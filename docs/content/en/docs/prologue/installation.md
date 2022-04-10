---
title: "Installation"
description: "Panduan ini berguna untuk membantu Anda menginstall Anchor CLI agar dapat digunakan di komputer Anda."
lead: "Panduan ini berguna untuk membantu Anda menginstall Anchor CLI agar dapat digunakan di komputer Anda."
date: 2020-10-06T08:48:57+00:00
lastmod: 2020-10-06T08:48:57+00:00
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 100
toc: true
---

## Installing Anchor

Setiap rilis Anchor menyediakan rilis biner untuk OS Linux dan Windows. Versi biner ini dapat diunduh dan diinstal secara manual. 

[📥 DOWNLOAD ( Release Anchor CLI )](https://github.com/cocatrip/anchor/releases/latest/)


### Linux

1. Download version [Anchor CLI](https://github.com/cocatrip/anchor/releases/latest/) atau wget url version Anchor CLI.

```bash
wget https://github.com/cocatrip/anchor/releases/download/v0.7/anchor-v0.7-linux-amd64.tar.gz
```

2. Unpack file yang telah di download (`tar -xvf anchor-v0.7-linux-amd64.tar.gz`)

3. Pindahkan binary Anchor CLI ke directory `bin` ( `/your-PATH/bin/anchor` atau `/your-PATH/local/bin/anchor` )

4. Pastikan Anchor telah ter-install dengan menggunakan perintah

```bash
anchor --help
```

### Windows

1. Download rilis [Anchor CLI](https://github.com/cocatrip/anchor/releases/latest/) untuk Windows

2. Extract file anchor-v0.7-windows-amd64.zip ( extract sesuai dengan lokasi yang diinginkan )

3. Input file anchor ke dalam Environment Variable Windows.

    a. Buka tab Windows → Search “Edit the system environment  variables”

    b. Buka Edit the system environment variables → buka Environtment Variables...

    ![env-var](https://fragrant-dessert-ea4.notion.site/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fdef24d17-58c7-4b85-ac4a-6d0771355d2e%2FUntitled.png?table=block&id=544300d5-efc9-4532-b46e-f9b525bb9495&spaceId=19beabee-e334-4147-ac83-9c7c735f986e&width=790&userId=&cache=v2)

    c. Input variable baru di User Variable for $USER dan System variables.

    d. Input dengan lokasi dari file anchor yang telah di extract.

    ![explorer1](https://fragrant-dessert-ea4.notion.site/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F3c186332-e497-436d-be8f-91f2e34f6a4e%2FUntitled.png?table=block&id=2d0968df-7b8e-412d-aae0-4cd078cbe716&spaceId=19beabee-e334-4147-ac83-9c7c735f986e&width=2000&userId=&cache=v2)

4. Pastikan anchor telah ter-install dengan membuka cmd dan ketikan perintah.

```bash
anchor --help
```
