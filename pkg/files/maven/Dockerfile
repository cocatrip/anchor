FROM trainingad1/java11
RUN java -version
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" > /etc/timezone
EXPOSE 1020
COPY [[ .Docker.JAR_SNAPSHOT ]].jar [[ .Docker.JAR_SNAPSHOT ]].jar
COPY . .
#RUN ls -la
RUN pwd
ENTRYPOINT ["java", "-javaagent:/[[ .Docker.JAR_ELASTIC ]].jar", "-Delastic.apm.service_name=[[ .Global.TESTING_TAG ]]-[[ .Global.APPLICATION_NAME ]]", "-Delastic.apm.server_urls=http://[[ .Docker.ELASTIC_URL ]]:[[ .Docker.ELASTIC_PORT ]]", "-Delastic.apm.application_packages=id.co.adira.[[ .Global.BUSINESS_NAME ]]" ,"-Xms512m" , "-Xmx1024m", "-jar", "[[ .Docker.JAR_SNAPSHOT ]].jar"]
#ENTRYPOINT java -Xms512m -Xmx1024m -jar [[ .Docker.JAR_SNAPSHOT ]].jar
