# Lambda RajaSMS Monitor
Monitor saldo credit dan tanggal kedaluarsa akun RajaSMS via discord dengan AWS Lambda

## Cara pakai

Get the things ready:

```bash
# Clone the repository
git clone git@github.com:xpartacvs/lambda-rajasms-monitor.git

# Change directory
cd lambda-rajasms-monitor

# Download the dependencies
go mod tidy

# Compile to binary
GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ./rajasms-monitor main.go

# Compress the binary
zip rajasms-monitor.zip rajasms-monitor
```

Lalu upload ke function AWS Lambda

## Konfigurasi

Konfigurasi aplikasi ini dapat dilakukan dengan menggunakan environment variables.

| **Variable**            | **Type**  | **Req** | **Default**             | **Description**                                                                                                                 |
| :---                    | :---      | :---:   | :---                    | :---                                                                                                                            |
| `DISCORD_WEBHOOKURL`    | `string`  | √       |                         | URL webhook Discord.                                                                                                            |
| `DISCORD_BOT_NAME`      | `string`  |         | suka-suka discord       | Nama bot yang akan muncul di channel Discord.                                                                                   |
| `DISCORD_BOT_AVATARURL` | `string`  |         | suka-suka discord       | URL ke file gambar yang akan digunakan sebagai avatar bot discord.                                                              |
| `DISCORD_BOT_MESSAGE`   | `string`  |         | `Reminder akun RajaSMS` | Pesan yang akan ditulis bot discord perihal status akun RajaSMS.                                                                |
| `LOGMODE`               | `string`  |         | `disabled`              | Mode log aplikasi: `debug`, `info`, `warn`, `error`, dan `disabled`.                                                            |
| `RAJASMS_API_URL`       | `string`  | √       |                         | URL server akun RajaSMS.                                                                                                        |
| `RAJASMS_API_KEY`       | `string`  | √       |                         | API key akun RajaSMS.                                                                                                           |
| `RAJASMS_LOWBALANCE`    | `integer` |         | `100000`                | Jika saldo <= nilai variabel ini maka alert via discord webhook akan terpicu.                                                   |
| `RAJASMS_GRACEPERIOD`   | `integer` |         | `7`                     | Jumlah hari menjelang tanggal kedaluarsa akun. Alert akan terpicu jika tanggal sekarang >= (tanggal kedaluarsa - variabel ini). |
