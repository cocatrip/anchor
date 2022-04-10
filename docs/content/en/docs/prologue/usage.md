---
title: "Usage"
description: "Panduan ini menjelaskan dasar-dasar dari penggunaan Anchor untuk pengelolaan konfigurasi secara otomatis. Disini Saya menganggap Anda telah menginstall anchor di panduan sebelumnya."
lead: ""
date: 2020-10-13T15:21:01+02:00
lastmod: 2020-10-13T15:21:01+02:00
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 130
toc: true
---

Panduan ini menjelaskan dasar-dasar dari penggunaan Anchor untuk pengelolaan konfigurasi secara otomatis. Disini Saya menganggap Anda telah menginstall anchor di
[panduan sebelumnya](/docs/installation/).

Jika Anda hanya tertarik untuk menjalankan beberapa perintah cepat, Anda dapat langsung beralih ke
[Panduan Memulai Cepat](/docs/quick-start/). Bab ini mencakup rincian perintah Anchor dan menjelaskan cara menggunakan Anchor

## Tiga Konfigurasi di Dalam Anchor 

Dockerfile, Dockerfile adalah dokumen teks yang berisi semua perintah yang dapat dipanggil pengguna pada baris perintah untuk membentuk image.
Dockerfile biasanya berisi sebuah perintah untuk membungkus sebuah aplikasi agar aplikasi tersebut dapat berjalan dalam sebuah container.
Proses membungkus Dockerfile menjadi sebuah image dapat dengan berbagai cara seperti melakukan build dengan perintah docker build atau dapat juga
menggunakan kaniko executor dan memasukannya ke dalam sebuah pipeline agar proses build dan yang lainnya dapat berjalan secara otomatis.

Jenkinsfile, Jenkinsfile adalah dokumen teks yang berisi perintah dalam bentuk pipeline. Jenkinsfile biasanya akan dimasukan ke dalam sebuah
repository yang nantinya dapat dipanggil oleh aplikasi Jenkinsfile menggunakan gitSCM.

Helm, Helm adalah alat penerapan Kubernetes untuk mengotomatiskan pembuatan, pengemasan, konfigurasi, dan penerapan aplikasi dan layanan ke kluster Kubernetes.

## 'anchor template' : generate configuration 

Anchor dapat meng-generate kebutuhan konfigurasi deployment Anda secara otomatis dengan:

- `anchor template docker`, dapat meng-generate Dockerfile
- `anchor template jenkins`, dapat meng-generate Jenkinsfile
- `anchor template helm`, dapat meng-generate Helm
- `anchor template all`, dapat meng-generate Dockerfile, Jenkinsfile dan Helm

Anda dapat membuat Dockerfile dengan perintah `anchor template docker`:

```bash
$ anchor template docker
Dockerfile-Default
------------------
FROM trainingad1/java11
RUN java -version
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" > /etc/timezone
EXPOSE 1020
COPY default-0.0.1-snapshot.jar default-0.0.1-snapshot.jar
COPY . .
#RUN ls -la
RUN pwd
ENTRYPOINT ["java", "-javaagent:/default-elastic-apm-agent-1.28.1.jar", "-Delastic.apm.service_name=Default-Default", "-Delastic.apm.server_urls=http://apm-server.logging.svc.cluster.local:8200", "-Delastic.apm.application_packages=id.co.adira.Default" ,"-Xms512m" , "-Xmx1024m", "-jar", "default-0.0.1-snapshot.jar"]
#ENTRYPOINT java -Xms512m -Xmx1024m -jar com.adira.leadengine.main-0.0.1-SNAPSHOT.jar
Parse SUCCESS!
```

Setelah Anda menjalankan perintah `anchor template docker` maka akan terbentuk 2 Dockerfile, yaitu:

1. Dockerfile-$TESTING_TAG
2. Dockerfile (template)

Dockerfile (template) berguna untuk memasukan beberapa `variable` agar dapat diparse langsung dengan `value` yang telah di set di `config.yaml`.

Example:

```dockerfile
FROM trainingad1/java11
RUN java -version
RUN cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" > /etc/timezone
EXPOSE 1020
# Set variable JAR_SNAPSHOT
COPY [[ .Docker.JAR_SNAPSHOT ]].jar [[ .Docker.JAR_SNAPSHOT ]].jar
COPY . .
#RUN ls -la
RUN pwd
# Set variable JAR_ELASTIC, TESTING_TAG, ELASTIC_URL, ELASTIC_PORT, BUSINESS_NAME, JAR_SNAPSHOT
ENTRYPOINT ["java", "-javaagent:/[[ .Docker.JAR_ELASTIC ]].jar", "-Delastic.apm.service_name=[[ .Global.TESTING_TAG ]]-[[.Global.APPLICATION_NAME]]", "-Delastic.apm.server_urls=http://[[.Docker.ELASTIC_URL]]:[[.Docker.ELASTIC_PORT]]", "-Delastic.apm.application_packages=id.co.adira.[[.Global.BUSINESS_NAME]]" ,"-Xms512m" , "-Xmx1024m", "-jar", "[[.Docker.JAR_SNAPSHOT]].jar"]
```

Dapat dilihat code diatas, beberapa variable di set di dalam Dockerfile (template) agar value dapat di parse langsung menggunakan `config.yaml`.

```yaml
docker:
  # Arbitrary / Optional / Dynamic
  JAR_SNAPSHOT: default-0.0.1-snapshot
  JAR_ELASTIC: default-elastic-apm-agent-1.28.1
  ELASTIC_URL: apm-server.logging.svc.cluster.local
  ELASTIC_PORT: 8080
```

Dari beberapa variable diatas, Anda harus me-set variable di dalam “[[ ]]” dan Anda harus menggunakan `.Global` untuk memanggil variable secara global di beberapa file berbeda dan `.Docker` `.Jenkins` `.Helm` untuk variable yang spesifik hanya dipanggil di file tersebut.

Kenapa perlu ada testing tag? perbedaan environment membuat testing tag diperlukan untuk membedakan jenis file yang akan digunakan disetiap proses yang berbeda-beda, sehingga di dalam anchor di define menggunakan testing tag untuk memisahkan proses QC deployment.

## '--config' atau '-c' : spesifik file config

Option `--config` atau `-c` digunakan untuk menargetkan lokasi config.yaml berada.

```bash
anchor template [template name] -c dir/config.yaml
```

## Membaca Teks Help

Untuk mempelajari lebih lanjut tentang perintah Anchor yang tersedia, ketik perintah diikuti dengan tanda `-h` atau `--help`:

```bash
anchor -h
```
