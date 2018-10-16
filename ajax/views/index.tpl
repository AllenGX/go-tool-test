<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Page Title</title>
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <script src="/static/js/jquery-3.3.1.min.js"></script>
</head>
<body>
  <h1 class="a bb ccc">ajax</h1>
  <script>
    var $h1=$('h1').text()
    var $class=$('h1').attr("class")
    console.log($h1)
    console.log($class)
    var $classQuenue=$class.split(" ")
    $classQuenue.forEach(element => {
      console.log(element)
    });


    $.ajax({ 
          type: "Get", 
          url:"http://127.0.0.1:8080/data",
          success:function(data){
            console.log(data.key)
          }
        }
      );

    console.log($.key);
var music="<ul>"; 
  </script>
</body>
</html>