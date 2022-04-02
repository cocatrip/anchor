pipeline {
  environment {
    Version_Major = 1
    Version_Minor = 0
    Version_Patch = 0
    IMAGE_NAME = "%{SERVER_NAME}/%{BUSINESS_NAME}/%{TESTING_TAG}-%{APPLICATION_NAME}"
    IMAGE_TAG = "${Version_Major}.${Version_Minor}.${Version_Patch}-${BUILD_TIMESTAMP}-${env.BUILD_NUMBER}"
    cluster_context = credentials('cluster-context')
    cluster_user = credentials('cluster-user')
    String SonarProjectKey = '%{APPLICATION_NAME}'
    String SonarHost = '%{LINK_SONARQUBE}'
    String SPRING_ACTIVE_PROFILE = '%{TESTING_TAG}'
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
            mvn  clean package 
          """
        }
      }
    }
    stage('sonarqube') {
      steps {
          withSonarQubeEnv('sonarqube-%{TESTING_TAG}') {
        container('maven') {
          sh """
           echo "******** currently executing sonarqube stage ********"
           mvn clean verify sonar:sonar sonar:sonar \
             -Dsonar.projectKey=%{APPLICATION_NAME} \
             -Dsonar.host.url=%{SONARQUBE_URL} \
             -Dsonar.login=%{SONARQUBE_KEY}
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
             cp $workspace/target/%{JAR_APP_NAME} ./%{JAR_APP_NAME}
           /kaniko/executor --dockerfile `pwd`/Dockerfile-%{TESTING_TAG} --context `pwd` --destination="${IMAGE_NAME}:${IMAGE_TAG}"
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
            helm upgrade -i ad1-lead-main helm/ad1-lead-main -f helm/ad1-lead-main/values-dev.yaml -n leadengine  --set=image.tag=${IMAGE_TAG}
            kubectl rollout status deployment/ad1-lead-main -n leadengine
            kubectl get pods -n leadengine
            helm ls -n leadengine
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
}