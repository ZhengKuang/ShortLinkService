# ShortURL 业务逻辑梳理
##用户第一次使用短链接服务
首先修改了配置文件
127.0.0.1 改成了www.go.com
首先输入www.go.com，通过nginx这个虚拟映射，随机进入了这三个端口，这算个端口的服务已经在本地通过go run main.go -port=：8080【8081】【8082】开启了
```sh
upstream go_server_pool{
        server 127.0.0.1:8080 weight=1;
        server 127.0.0.1:8081 weight=2;
        server 127.0.0.1:8083 weight=1;
}

server {
        server_name www.go.com;
        charset utf-8;
        location ~* / {
           proxy_pass http://go_server_pool;
        }
}

```
之后用户登陆到了www.go.com的主界面，在输入栏中写入对应的长链接，会调用到index.html的javascrpt小函数
```sh
        <script type="text/javascript">
          $('#click').click(function(){
            empty();
            var sourceUrl=$('input[name=source]').val();
            $.get('/api/url',{'sourceUrl':sourceUrl},function(data){
              if(data.success){
                  var shortUrl=window.location.host+"/"+data.result.ShortUrl
                  console.log(window.location.host);
                  console.log(data.result.ShortUrl);
                  $('#result').html(shortUrl);
                  $('#result').attr("href","http://"+shortUrl);
                }else{
                  $('#message').html(data.result);
                }
            })

          })

            function empty(){
              $('#result').html("");
              $('#message').html("");
            }

          
        </script>
```
在这个函数中，用户以点击click，就会重新get一个链接，调用到
```sh
beego.RESTRouter("api/url", &controllers.UrlController{})
```
的UrlContorller，从而使用UrlController的Get方法
```sh
func (this *UrlController) Get() {
	sourceUrl := this.Ctx.Input.Query("sourceUrl")
	_, err := url.ParseRequestURI(sourceUrl)
	if err != nil {
		this.error("Url is not a valid url:" + err.Error())
		return
	}
	u := &models.Url{
		SourceUrl: sourceUrl,
	}
	u.GenId()
	u.ShortUrl = utils.IdToString(u.Id)
	u.Save()
	this.success(u)
}
```
在点击之后，会完成一系列的自增id，短链接形成，储存只在mongoDB中并且返回data数据，传回index.html中（相当于一个回调函数callback function），并且将相应的结果在index.html中显示出来。
##用户第二次使用短链接服务
比如说用户点击了输入了www.go.com/6这个短链接，会调用main.go里面的一行代码
```sh
	beego.Router("/*", &controllers.RedirectController{})
```
redirectionController里面会get到urlpath，并且拿到短链接6，并且用utils里面的StringToId找到自增id，通过FindByid方法返回对应的model.url结构{id，shorURL，SourceURL}，并且调用go的Redirct方法，返回302，并且返回SourceURL，在此，常规的业务逻辑已经完成了。
```sh
func (this *RedirectController) Get() {
	if this.Ctx.Request.URL.Path == "/" {
		fmt.Println("RediretController has been used, the url is ", this.Ctx.Request.URL.Path)
		this.Data["title_name"] = "short-link service_beego"
		this.TplName = "index.html"
		return

	}
	urlPath := this.Ctx.Request.URL.Path
	urlPath = strings.Trim(urlPath, "/")
	id := utils.StringToId(urlPath)
	u := &models.Url{
		Id: id,
	}
	err := u.FindById()
	if err != nil {
		this.error("no corresponding short URL in MongoDB:" + err.Error())
		return
	}
	this.Redirect(u.SourceUrl, 302)
}
```

总结一下，两个业务逻辑总共运用到的东西是：1.nginx监听www.go.com并且返回到对应的本地8080，8081，8082端口。（调用到main.go和redirec.go）2.用户输入长链接，这时候重新发出一个get请求，其请求api/url目录下的东西，这时候main刚好在有对应的controller，所以在这个controller里面执行了自增id，用自增id生成短链接，储存到mongodb中，并且返回success的data（调用到main.go,以及controller里面的url.go）3.用户点击短链接，比如www.go.com/6这个链接，会被main再次监听到，然后调用redirectController的方法，找到id，映射成长链接并且返回（调用到main.go以及controller里面的redirec.go）



