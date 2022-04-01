FROM trainingad1/java11
RUN java -version
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" > /etc/timezone
EXPOSE 1020
COPY %{JAR_SNAPSHOT}.jar %{JAR_SNAPSHOT}.jar
COPY . .
#RUN ls -la
RUN pwd
ENTRYPOINT ["java", "-javaagent:/%{JAR_ELASTIC}.jar", "-Delastic.apm.service_name=%{TESTING_TAG}-%{APPLICATION_NAME}", "-Delastic.apm.server_urls=http://%{ELASTIC_URL}:%{ELASTIC_PORT}", "-Delastic.apm.application_packages=id.co.adira.%{BUSINESS_NAME}" ,"-Xms512m" , "-Xmx1024m", "-jar", "%{JAR_SNAPSHOT}.jar"]
#ENTRYPOINT java -Xms512m -Xmx1024m -jar com.adira.leadengine.main-0.0.1-SNAPSHOT.jar
