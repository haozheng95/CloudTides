const express = require('express')
const app = express()
const ejs = require('ejs')
const path = require('path')
const port = 3000
app.use(express.static(path.join(__dirname, 'tides-ui')));
app.set("views", __dirname + "/tides-ui");
app.engine('html',ejs.__express)
// 配置模版引擎
app.set("view engine", "html")
app.get('/', (req, res) => {
  let resp = {}
  // const data = fs.readFileSync(path.resolve(__dirname,'\dist\index.html'));
  // fs.readFile(path.resolve(__dirname,'\views','\index.html'), function (err, data) {
  //   if (err) {
  //       return console.error(err);
  //   }
  //   resp = data
  //   console.log("异步读取: " + data.toString());
  // });
  res.render('index', {})
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
