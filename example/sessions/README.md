gorilla/sessions为自定义session后端提供cookie和文件系统session以及基础结构。

主要功能是：

* 简单的API：将其用作设置签名（以及可选的加密）cookie的简便方法。
* 内置的后端可将session存储在cookie或文件系统中。
* Flash消息：一直持续读取的session值。
* 切换session持久性（又称“记住我”）和设置其他属性的便捷方法。
* 旋转身份验证和加密密钥的机制。
* 每个请求有多个session，即使使用不同的后端也是如此。
* 自定义session后端的接口和基础结构：可以使用通用API检索并批量保存来自不同商店的session。

