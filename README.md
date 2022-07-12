# Moba-Xterm

## 介绍
    https://github.com/luckylocode/mobaXterm 为 Java 移植版,版本要求 Java 1.8 ,基于 JavaFX 开发 GUI;
    自从 Java 9 开始, JavaFX 从 JDK 中剔除,需要单独安装,为避免版本升级导致JDK(JRE)没有 JavaFX API 的问题,
    该工具用 Golang 重写.
    
    Golang 没有官方的 UI 库, 本仓库 选用 fyne 这个库开发, 使用 fyne-cross 交叉编译.
    
    原作者 GitHub Python 版地址为:
    https://github.com/DoubleLabyrinth/MobaXterm-keygen.git

## 编译说明
    1 : 打包资源文件 请使用  linux 的 shell  或者 gitbash, 参见 bundled.sh
    2 : 编译需要 docker . 务必备好 科学上网.
    3 : linux 未尝试, 本编译基于  Windows10(11) -- wsl2 -- docker

### wsl2 环境 准备
    1: windows 安装 https://docs.microsoft.com/zh-cn/windows/wsl/install-manual   
    
    2: 基于 wsl2 安装docker,推荐 第一种方式
       https://zhuanlan.zhihu.com/p/421998834
    
    3: 安装 golang 
       wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
       rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz
       export PATH=$PATH:/usr/local/go/bin
       go env -w GO111MODULE=on
       go env -w GOPROXY=https://goproxy.cn,direct

    4: 源码下载 (或者自行下载源码包)
       git clone https://github.com/luckylocode/Moba-Xterm.git

### 网络准备
    1: 宿主机允许 科学上网允许来自局域网的连接
    
    2: wsl2 配置网络代理,走宿主机的 科学上网(以 v2ray 为例; 科学上网允许来自局域网的连接 !!!),
    vi ~/.bashrc 追加如下配置

```shell
export windows_host=`cat /etc/resolv.conf|grep nameserver|awk '{print $2}'`
export ALL_PROXY=socks5://$windows_host:10808

export HTTP_PROXY=$ALL_PROXY
export http_proxy-$ALL_PROXY
export HTTPS_PROXY=$ALL_PROXY
export https_proxy=$ALL_PROXY

if [ "`git config --global --get proxy.https`" != "socks5://$windows_host:10808" ]; then
            git config --global proxy.https socks5://$windows_host:10808
fi
```
    3: docker 运行时代理
       fyne-cross 交叉编译需要 docker ,编译时 运行的容器会去拉取相关依赖,需要科学上网
       vi  ~/.docker/config.json; windows_host 参见上一步获取的 ip

```json
{
 "proxies":
 {
   "default":
   {
     "httpProxy": "http://windows_host:10808",
     "httpsProxy": "http://windows_host:10808",
     "noProxy": "*.test.example.com,.example2.com,127.0.0.0/8"
   }
 }
}
```
