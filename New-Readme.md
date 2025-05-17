# Ushio Prober

A maimai DX prober based on [Diving-Fish/maimaidx-prober](https://github.com/diving-fish/maimaidx-prober), with new user experience by new user interface.

## User Guide

*You will need a account to use our service. This account will record your score data every time you upload them by any method.*

If you encountered any problem when using our service, please check the F&Q at the end of this file. If F&Q can not resolve your problem, feel free to open an issue to report the problem.

## How to import your data?

There are several way that you can upload your score data to our service.

### 1. Import your data via Proxy Tool

Thanks the people who created and contributed to the proxy tool:
- [Diving-Fish](https://github.com/Diving-Fish)
- [xen0n](https://github.com/xen0n)
- yixin.yu(No link available)
- [wzhqwq](https://github.com/wzhqwq)

#### Preparations

You can open [This link](https://github.com/Ushio-Prober/releases) to download the Proxy Tool, and we will tell you which version of Proxy Tool you need to download decided to which operating system you are using.

Please copy the Proxy Tool to an empty folder, and double-click it or run in a terminal. When you first running this tool, it will create 3 files in current folder:

|File name|Description|
|------|----|
|`config.json`|Configuration file of the Proxy Tool, you will need to write the username and password there.|
|`cert.crt`|A certification file you can import to your os.|
|`key.pem`|Private key file. Any operation is not needed.|

In Windows, you need to double-click the `cert.crt` file, click the **Install Cert**, open the cert import guide. At most time, you only want to install the cert file for current user.

In next step, you will need choose to put the cert file into the "Trusted Root Certification Authorities" folder to storage the cert file. Click "OK" for the remaining steps until Windows tells you that cert file imported successfully.

![](https://www.diving-fish.com/images/maimaidx-prober/1.png)

<details>

<summary>Setup for Unix-like series operating system:</summary>
&nbsp;

*If you're running the Proxy Tool under a macOS or GNU/Linux platform, type the commands lists below in terminal:*

```plain
chmod +x <Filename>
./Filename>
```

*macOS users may encounter the "Untrusted Developer" problem. Please search for the solution in your own hands.*

*On macOS, please add the Keychain into the "System", and double-click the cert file to open the details, modify the settings under the "Trust" menu. Like this:*

![](https://www.diving-fish.com/images/maimaidx-prober/9.png)

</details>
&nbsp;

After done the step above, open the `config.json` file via a text editor. Write your username and password into the corresponding quotes like this:

```json
{"username": "MyAccount", "password": "MyPassword"}
```

Save the file, then open the Proxy Tool, the console may output the log like this:

```plain
2023/07/04 16:55:06 INFO: 您使用的是最新版本。
2023/07/04 16:55:06 INFO: 登录成功
2023/07/04 16:55:06 INFO: 使用此软件则表示您同意共享您在微信公众号舞萌 DX、中二节奏中的数据。
2023/07/04 16:55:06 INFO: 您可以在微信客户端访问微信公众号舞萌 DX、中二节奏的个人信息主页进行分数导入，如需退出请直接关闭程序或按下 Ctrl + C
2023/07/04 16:55:06 INFO: 代理设置已自动修改。
2023/07/04 16:55:06 INFO: 代理已开启到 127.0.0.1:8033
```

