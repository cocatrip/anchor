---
title: "Quick Start"
description: "Panduan ini berguna untuk Anda agar dapat cepat mulai menggunakan Anchor"
lead: "Panduan ini berguna untuk Anda agar dapat cepat mulai menggunakan Anchor"
date: 2020-11-16T13:59:39+01:00
lastmod: 2020-11-16T13:59:39+01:00
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 110
toc: true
---

## Requirements

Prasyarat berikut diperlukan untuk dapat menggunakan Anchor CLI dengan aman dan berhasil.

- Sebuah Aplikasi yang membutuhkan konfigurasi Dockerfile, Jenkinsfile dan Helm
- Helm CLI
- Anchor CLI

### Menyiapkan Aplikasi

Sebelum Anda memulai menggunakan Anchor CLI, Anda harus memastikan terlebih dahulu Anda memiliki sebuah Aplikasi yang membutuhkan beberapa konfigurasi berikut.

- Dockerfile
- Jenkinsfile
- Helm

### Install Anchor CLI

Download rilis biner dari klien Anchor CLI. Anda dapat melihat rilis biner di halaman [release](https://github.com/cocatrip/anchor/releases/latest/)

Kunjungi panduan menginstall Anchor CLI Untuk detail lebih lanjut, atau opsi lain. [Installation →](/docs/prologue/installation/)

## Membuat config.yaml

Agar anda dapat menggunakan Anchor CLI, Anda harus menyiapkan file config.yaml terlebih dahulu dengan membuatnya menggunakan text editor ( vim, nano, notepad, dsb ).

```yaml
### REQUIRED ###
APPLICATION_NAME: ad1-lead-main
BUSINESS_NAME: administrasi
SERVER_NAME: https://adira.co.id
TEMPLATE: maven
TESTING_TAG: uat

jenkins:
  # Arbitrary / Dynamic
  CREDENTIAL_HELM: kubeconfig
  JAR_APP_NAME: com.admengine-0.0.1-SNAPSHOT.jar
  LINK_SONARQUBE: http://10.161.17.144:9000/dashboard?id=
  SONARQUBE_KEY: alksfhoqhqhfijv0w9320582358flkshal
  SONARQUBE_URL: http://10.50.0.2:9000

docker:
  # Arbitrary / Dynamic
  ELASTIC_PORT: 8200
  ELASTIC_URL: apm-server.logging.svc.cluster.local
  JAR_ELASTIC: elastic-apm-agent-1.28.1
  JAR_SNAPSHOT: id-0.0.1-snapshot

helm:
  # Arbitrary / Dynamic
  BUILD_NUMBER: 1
  BUILD_TIMESTAMP: 220331
  CPU_LIMIT: 2500m
  CPU_REQUEST: 250m
  DB_PASSWORD: root
  DB_USER: root
  MEM_LIMIT: 2048Mi
  MEM_REQUEST: 256Mi
  NAMESPACE: leadengine
  SECRETDB_NAME: regcred
  SERVICE_PORT: 443
  SERVICE_TARGETPORT: 8081
  SERVICE_TYPE: LoadBalancer
  Version_Major: 1
  Version_Minor: 0
  Version_Patch: 0
```

config.yaml diatas merupakan default Anchor CLI, sehingga Anda dapat memulainya dengan cepat.

## Generate Project

Anda dapat generate project secara keseluruhan setelah Anda membuat file config.yaml dengan perintah.

```bash
$ anchor template all
Dockerfile  Dockerfile-uat  Jenkinsfile  Jenkinsfile-uat  config.yaml  helm
```

File Dockerfile, Jenkinsfile dan juga values.yaml yang tidak menggunakan $TESTING_TAG merupakan file template untuk mengubah template sesuai dengan template yang diinginkan. Untuk Helm, Helm yang digunakan sebagai default merupakan helm dengan konfigurasi menggunakan kubernetes secret dan configmap.

Jika Anda hanya ingin generate file yang diinginkan anda dapat menggunakan beberapa perintah berikut.

```bash
anchor template docker
anchor template jenkins
anchor template helm
anchor template helm --no-secret
```

## Membaca Teks Help

Untuk mempelajari lebih lanjut tentang perintah Anchor yang tersedia, ketik perintah diikuti dengan tanda -h atau --help:

```bash
anchor -h
```
