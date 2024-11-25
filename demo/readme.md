# demo代码

用于验证区块链

windows 环境

## goethereumbook 相关代码


```

geth相关包的使用

通过go代码对以太坊区块链进行操作


```

## newchain 新链数据

```
编译项目

新建账户
.\build\bin\geth.exe --datadir .\demo\newchain\ account new
密码统一为
a123456

Public address of the key:   0xbF2B6788585750cF62ae75755Ae9AFeE34828EC7
Path of the secret key file: demo\newchain\keystore\UTC--2024-11-25T09-00-02.614984800Z--bf2b6788585750cf62ae75755ae9afee34828ec7

Public address of the key:   0x015C8a48429584Bb535776Ed7691Ec721273D2f4
Path of the secret key file: demo\newchain\keystore\UTC--2024-11-25T09-01-00.014762500Z--015c8a48429584bb535776ed7691ec721273d2f4


初始化创世区块
.\build\bin\geth.exe --datadir .\demo\newchain\ init .\demo\genesis.json

启动私有区块链

.\build\bin\geth.exe --maxpeers 0 --datadir .\demo\newchain\ console
???

启动dev模式的区块链

.\build\bin\geth.exe --dev --datadir .\demo\devchain\ console


```