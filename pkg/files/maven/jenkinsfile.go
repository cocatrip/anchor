package maven

var Jenkinsfile string = `pipeline {
  environment {
    Version_Major = 1
    Version_Minor = 0
    Version_Patch = 0
    IMAGE_NAME = "[[ .Global.SERVER_NAME ]]/[[ .Global.BUSINESS_NAME ]]/[[ .Global.TESTING_TAG ]]-[[ .Global.APPLICATION_NAME ]]"
    IMAGE_TAG = "${Version_Major}.${Version_Minor}.${Version_Patch}-${BUILD_TIMESTAMP}-${env.BUILD_NUMBER}"
    cluster_context = credentials('cluster-context')
    cluster_user = credentials('cluster-user')
    String SonarProjectKey = '[[ .Global.APPLICATION_NAME ]]'
    String SonarHost = '[[ .Jenkins.LINK_SONARQUBE ]]'
    String SPRING_ACTIVE_PROFILE = '[[ .Global.TESTING_TAG ]]'
  }

  agent {
    kubernetes {
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
metadata:
labels:
  component: ci
spec:
  # Use service account that can deploy to all namespaces
  
  containers: 
  - name: maven
    image: maven:3.8.1-openjdk-11
    imagePullPolicy: IfNotPresent
    command:
    - cat
    tty: true 
  - name: helm
    image: trainingad1/helm3
    imagePullPolicy: IfNotPresent
    command:
    - cat
    tty: true
  - name: kaniko
    image: gcr.io/kaniko-project/executor:debug-539ddefcae3fd6b411a95982a830d987f4214251
    imagePullPolicy: IfNotPresent
    command:
    - /busybox/cat
    tty: true
    volumeMounts:
      - name: docker-config
        mountPath: /kaniko/.docker
  volumes:
  - name: docker-config
    projected:
      sources:
      - secret:
          name: regcred1
          items:
            - key: .dockerconfigjson
              path: config.json
    
      """
    }
  }

  stages {
    stage('build') {
      steps {
        container('maven') {
          sh """
            echo "******** currently executing Build stage ********"
            [[- if eq .Global.TEMPLATE "maven" ]]
            mvn clean package 
            [[- end ]]
            [[- if eq .Global.TEMPLATE "node" ]]
            npm install
						npm run build
            [[- end ]]
          """
        }
      }
    }

    stage('sonarqube') {
      steps {
        withSonarQubeEnv('sonarqube-[[ .Global.TESTING_TAG ]]') {
          container('maven') {
            sh """
              echo "******** currently executing sonarqube stage ********"
              mvn clean verify sonar:sonar sonar:sonar \
                -Dsonar.projectKey=[[ .Global.APPLICATION_NAME ]] \
                -Dsonar.host.url=[[ .Jenkins.SONARQUBE_URL ]] \
                -Dsonar.login=[[ .Jenkins.SONARQUBE_KEY ]]
            """
          }
        }
      }
    }

    stage("Quality gate") {
      steps {
        echo "******** currently executing Quality gate stage ********"
        waitForQualityGate abortPipeline: false
      }
    }

    stage('kaniko stage and pushing image to ali private registry') {
      steps {
        container('kaniko') {
          sh """
            echo "******** currently executing kaniko stage ********"
            cp $workspace/target/[[ .Jenkins.JAR_APP_NAME ]] ./[[ .Jenkins.JAR_APP_NAME ]]
            /kaniko/executor --dockerfile `+"`"+`pwd`+"`"+`/Dockerfile-[[ .Global.TESTING_TAG ]] --context `+"`"+`pwd`+"`"+` --destination="${IMAGE_NAME}:${IMAGE_TAG}"
          """
        }
      }
    }

    stage('helm') {
      steps {
        container('helm') {
          withCredentials([file(credentialsId: 'kubeconfig', variable: 'KUBECONFIG')]) {
            sh """
              echo "******** currently executing deployment stage ********"
              kubectl config set-context ${cluster_context} --cluster=kubernetes --user=${cluster_user}
              kubectl config use-context ${cluster_context}
              kubectl get nodes
              helm upgrade -i [[ .Global.APPLICATION_NAME ]] helm/[[ .Global.APPLICATION_NAME ]] -f helm/[[ .Global.APPLICATION_NAME ]]/values-[[ .Global.TESTING_TAG ]].yaml -n [[ .Jenkins.NAMESPACE ]]  --set=image.tag=${IMAGE_TAG}
              kubectl rollout status deployment/[[ .Global.APPLICATION_NAME ]] -n [[ .Jenkins.NAMESPACE ]]
              kubectl get pods -n [[ .Jenkins.NAMESPACE ]]
              helm ls -n [[ .Jenkins.NAMESPACE ]]
            """
          }
        }
      }
      post {
        failure {
          mail to: 'v.hanif.alfath@adira.co.id,nika.aditia@adira.co.id,lucky.andriawan@adira.co.id,sonarandkatalon@adira.co.id',
          subject: "FAILED: Build ${env.JOB_NAME}",
          body: "Build failed ${env.JOB_NAME} build no: ${env.BUILD_NUMBER}.\n\nView the log at:\n ${env.BUILD_URL}\n\nSee The SonarQube Report : ${SonarHost}${SonarProjectKey}"
        }

        success{
          mail to: 'v.hanif.alfath@adira.co.id,nika.aditia@adira.co.id,lucky.andriawan@adira.co.id,sonarandkatalon@adira.co.id',
          subject: "SUCCESSFUL: Build ${env.JOB_NAME}",
          body: "Build Successful ${env.JOB_NAME} build no: ${env.BUILD_NUMBER}\n\nView the log at:\n ${env.BUILD_URL}\n\nSee The SonarQube Report : ${SonarHost}${SonarProjectKey}"
        }

        aborted{
          mail to: 'v.hanif.alfath@adira.co.id,nika.aditia@adira.co.id,lucky.andriawan@adira.co.id,sonarandkatalon@adira.co.id',
          subject: "ABORTED: Build ${env.JOB_NAME}",
          body: "Build was aborted ${env.JOB_NAME} build no: ${env.BUILD_NUMBER}\n\nView the log at:\n ${env.BUILD_URL}\n\nSee The SonarQube Report : ${SonarHost}${SonarProjectKey}"
        }
      } 
    }
  }
}`
