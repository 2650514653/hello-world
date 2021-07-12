<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Document</title>
</head>
<body>
  <h2>1.模板中绑定基本的数据 字符串 数值 布尔值</h2>
  <p>{{.title}}</p>
  <p>{{.num}}</p>
  <p>{{.flag}}</p>

  <h2>1.模板中绑定结构体</h2>
  <p>{{.article.Title}}</p> 
  <p>{{.article.Content}}</p>

  <h2>模板中自定义变量</h2>
  <p>{{$title := .title}}</p>
  <p>{{$title}}</p>

  <h2>模板中循环遍历切片</h2>
  <ul>
    {{range $key,$value:= .sliceList }}
    <li>{{$key}}---{{$value}}</li>
    {{ end }}  
  </ul>
 
  <ul>
    <li>

    </li>
  </ul>
  <p></p>

</body>
</html>