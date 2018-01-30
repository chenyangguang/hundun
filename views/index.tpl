<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="context-type"  content="text/html"; charset="utf-8">
  <title>Go-beego-vue</title>

  <link rel="stylesheet" type="text/css" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.6.3/css/font-awesome.min.css">

  <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/jquery/2.2.4/jquery.min.js"></script>
  <script type="text/javascript" src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js"></script>
  <script type="text/javascript" src="http://cdnjs.cloudflare.com/ajax/libs/vue/1.0.24/vue.min.js"></script>
  <script type="text/javascript" src="https://cdnjs.cloudflare.com/ajax/libs/vue-resource/0.7.0/vue-resource.min.js"></script>
</head>
<body>

    <div class="container">
        <div class="row">
          <div class="col-md-6">
            <h1>My Stock</h1>

            <div class="input-group">
              <input type="text" class="form-control" placeholder="New stock" v-on:keyup.enter="createTask" v-model="newTask.name" autofocus>
              <span class="input-group-btn">
                  <button class="btn btn-primary" type="button" v-on:click="Search">Search</button>
              </span>
            </div>

          </div>
        </div>
    </div>
<script type="text/javascript" src="static/js/app.js"></script>
</body>
</html>
