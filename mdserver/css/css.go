package css

// Stylesheets
//This file is autogenerated from assets directory
const (
	CSS = `@font-face{font-family:FiraCode-Regular;src:url(https://unpkg.com/firacode/distr/ttf/FiraCode-Regular.ttf)}.switch{position:relative;display:inline-block;width:60px;height:34px}.switch input{opacity:0;width:0;height:0}.slider{position:absolute;cursor:pointer;top:0;left:0;right:0;bottom:0;background-color:#f8efe1;-webkit-transition:.4s;transition:.4s}.slider:before{position:absolute;content:"";height:26px;width:26px;left:4px;bottom:4px;background-color:#1c1c1c;-webkit-transition:.4s;transition:.4s}.slider.round{border-radius:34px}.slider.round:before{border-radius:50%}input:checked+.slider:before{background-color:#f8efe1}input:checked+.slider{background-color:#1c1c1c}input:checked+.slider:before{background-color:#f8efe1;-webkit-transform:translateX(26px);-ms-transform:translateX(26px);transform:translateX(26px)}input:focus+.slider{box-shadow:0 0 1px #1c1c1c}.top-bar{display:flex}.dropbtn{padding:16px;font-size:16px;border:0}.dropdown{position:relative;display:inline-block}.dropdown:hover .dropdown-content{display:block}.dropdown-content{display:none;position:absolute;min-width:160px;box-shadow:0 8px 16px 0 rgba(0,0,0,.2);z-index:1}.dropdown-content a{padding:12px 16px;text-decoration:none;display:block}a{text-decoration:none}a.anchor{display:block;padding-left:30px;margin-left:-30px;cursor:pointer;position:absolute;top:0;left:0;bottom:0}:root{--line-no-width:2em}pre{white-space:pre-wrap}pre:before{counter-reset:listing}pre code{counter-increment:listing}pre code:before{content:counter(listing) ". ";display:inline-block;width:var(--line-no-width);padding-left:auto;margin-left:auto;text-align:right;font-weight:590}`

	FvCommon = `*{margin:0;padding:0}body{display:flex;flex-direction:column;height:100vh;padding-top:30px}nav{margin:50px}ul{list-style-type:none}li{height:25px;margin-right:0;padding:0 20px}li.active a{font-weight:700}li a{text-decoration:none;font:25px/1 FiraCode-Regular,monospace,sans-serif;-webkit-transition:all .5s ease;-moz-transition:all .5s ease;-o-transition:all .5s ease;-ms-transition:all .5s ease;transition:all .5s ease}.files{margin:auto;height:50%;padding:10px;display:flex;flex-direction:column}.top-bar{display:flex;padding-left:30px;line-height:1.6;width:80%;margin:0 auto}`

	FvDark = `.bbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #666;color:#f8efe1;font-size:25px;line-height:1.2em}.bbtn:hover{color:#000;background-color:#f8efe1}.tbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #666;color:#f8efe1;font-size:14px;line-height:2.2em;width:auto}.tbtn:hover{color:#000;background-color:#f8efe1}.tbtn a{color:#f8efe1}.dropdown-content{background-color:#f8efe1}.dropdown-content a{color:#1c1c1c}.dropdown-content a:hover{background-color:#ccc}body{background:#1c1c1c}li.active a{color:#333}li a{color:#ccc}li a:hover{color:#666}`

	FvLight = `.bbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #1c1c1c;color:#1c1c1c;font-size:25px;line-height:1.2em}.bbtn:hover{color:#ccc;background-color:#1c1c1c}.tbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #1c1c1c;color:#1c1c1c;font-size:14px;line-height:2.2em;width:auto}.tbtn:hover{color:#ccc;background-color:#1c1c1c}.tbtn a{color:#1c1c1c}.dropdown-content{background-color:#1c1c1c}.dropdown-content a{color:#fff}.dropdown-content a:hover{background-color:#333}body{background:#f8efe1}li.active a{color:#333}li a{color:#ccc}li a:hover{color:#666}`

	Ghmd = `.sidebar{height:100%;width:0;position:fixed;z-index:1;top:0;left:0;background-color:#f8efe1;overflow-x:hidden;transition:.5s;padding-top:60px}.sidebar a{padding:8px 8px 8px 32px;text-decoration:none;font-size:25px;color:#666;display:block;transition:.3s;font-size:18px}.sidebar a:hover{color:#333}.sidebar .closebtn{position:absolute;top:0;right:25px;font-size:36px;margin-left:50px}#main{transition:margin-left .5s;padding:16px}@media screen and (max-height:450px){.sidebar{padding-top:15px}}body{font-family:FiraCode-Regular,monospace,sans-serif;font-size:14px;line-height:1.6;padding-top:10px;padding-bottom:10px;padding:30px}body>:first-child{margin-top:0!important}body>:last-child{margin-bottom:0!important}body>h2:first-child{margin-top:0;padding-top:0}body>h1:first-child{margin-top:0;padding-top:0}body>h1:first-child+h2{margin-top:0;padding-top:0}a{color:#4183c4;text-decoration:none}a.absent{color:#c00}a.anchor{display:block;padding-left:30px;margin-left:-30px;cursor:pointer;position:absolute;top:0;left:0;bottom:0}h1,h2,h3,h4,h5,h6{margin:20px 0 10px;padding:0;font-weight:700;-webkit-font-smoothing:antialiased;cursor:text;position:relative}h1:first-child,h1:first-child+h2,h2:first-child,h3:first-child,h4:first-child,h5:first-child,h6:first-child{margin-top:0;padding-top:0}h1:hover a.anchor,h2:hover a.anchor,h3:hover a.anchor,h4:hover a.anchor,h5:hover a.anchor,h6:hover a.anchor{text-decoration:none}h1 code,h1 tt{font-size:inherit}h2 code,h2 tt{font-size:inherit}h3 code,h3 tt{font-size:inherit}h4 code,h4 tt{font-size:inherit}h5 code,h5 tt{font-size:inherit}h6 code,h6 tt{font-size:inherit}h1{font-size:28px}h2{font-size:24px;border-bottom:1px solid #ccc}h3{font-size:18px}h4{font-size:16px}h5{font-size:14px}h6{color:#777;font-size:14px}blockquote,dl,li,ol,p,pre,table,ul{margin:15px 0}hr{border:0 none;color:#ccc;height:4px;padding:0}body>h3:first-child,body>h4:first-child,body>h5:first-child,body>h6:first-child{margin-top:0;padding-top:0}a:first-child h1,a:first-child h2,a:first-child h3,a:first-child h4,a:first-child h5,a:first-child h6{margin-top:0;padding-top:0}h1 p,h2 p,h3 p,h4 p,h5 p,h6 p{margin-top:0}li p.first{display:inline-block}ol,ul{padding-left:30px}ol :first-child,ul :first-child{margin-top:0}ol :last-child,ul :last-child{margin-bottom:0}dl{padding:0}dl dt{font-size:14px;font-weight:700;font-style:italic;padding:0;margin:15px 0 5px}dl dt:first-child{padding:0}dl dt>:first-child{margin-top:0}dl dt>:last-child{margin-bottom:0}dl dd{margin:0 0 15px;padding:0 15px}dl dd>:first-child{margin-top:0}dl dd>:last-child{margin-bottom:0}blockquote{border-left:4px solid #ddd;padding:0 15px;color:#777}blockquote>:first-child{margin-top:0}blockquote>:last-child{margin-bottom:0}table{padding:0}table tr{border-top:1px solid #ccc;background-color:#fff;margin:0;padding:0}table tr:nth-child(2n){background-color:#f8f8f8}table tr th{font-weight:700;border:1px solid #ccc;text-align:left;margin:0;padding:6px 13px}table tr td{border:1px solid #ccc;text-align:left;margin:0;padding:6px 13px}table tr td :first-child,table tr th :first-child{margin-top:0}table tr td :last-child,table tr th :last-child{margin-bottom:0}img{display:block;max-height:100%;max-width:100%}span.frame{display:block;overflow:hidden}span.frame>span{border:1px solid #ddd;display:block;float:left;overflow:hidden;margin:13px 0 0;padding:7px;width:auto}span.frame span img{display:block;float:left}span.frame span span{clear:both;color:#333;display:block;padding:5px 0 0}span.align-center{display:block;overflow:hidden;clear:both}span.align-center>span{display:block;overflow:hidden;margin:13px auto 0;text-align:center}span.align-center span img{margin:0 auto;text-align:center}span.align-right{display:block;overflow:hidden;clear:both}span.align-right>span{display:block;overflow:hidden;margin:13px 0 0;text-align:right}span.align-right span img{margin:0;text-align:right}span.float-left{display:block;margin-right:13px;overflow:hidden;float:left}span.float-left span{margin:13px 0 0}span.float-right{display:block;margin-left:13px;overflow:hidden;float:right}span.float-right>span{display:block;overflow:hidden;margin:13px auto 0;text-align:right}code,tt{margin:0 2px;padding:0 5px;white-space:nowrap;border-radius:3px}pre{font-family:FiraCode-Regular,monospace,sans-serif;font-size:14px;line-height:19px;overflow:auto;padding:6px 10px;border-radius:3px}pre code{margin:0;padding:0;white-space:pre;border:0;background:0 0}.highlight pre{font-size:13px;line-height:19px;overflow:auto;padding:6px 10px;border-radius:3px}pre code,pre tt{background-color:transparent;border:0}.theme-switch{margin:0 5px}@media all and (max-width:30em){a.bbtn{display:block;margin:.4em auto}}@media all and (min-width:1300px){.wrapper{width:80%;margin:0 auto}}`

	GhmdDark = `.bbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #666;color:#f8efe1;font-size:25px;line-height:1.2em}.bbtn:hover{color:#000;background-color:#f8efe1}.tbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #666;color:#f8efe1;font-size:14px;line-height:2.2em;width:auto}.tbtn:hover{color:#000;background-color:#f8efe1}.tbtn a{color:#f8efe1}.dropdown-content{background-color:#f8efe1}.dropdown-content a{color:#1c1c1c}.dropdown-content a:hover{background-color:#ccc}body{background-color:#1c1c1c;color:#f8efe1}h1,h2,h3,h4,h5,h6{color:#f8efe1}code{background-color:#333;border:1px solid #666}pre{color:#f8efe1;background-color:#232323;border:1px solid #333}pre.highlight{background-color:#f8efe1;border:1px solid #ccc}.token-add{background:green}.token-del{background:#c00000}`

	GhmdLight = `.bbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #1c1c1c;color:#1c1c1c;font-size:25px;line-height:1.2em}.bbtn:hover{color:#ccc;background-color:#1c1c1c}.tbtn{display:inline-block;box-sizing:border-box;text-decoration:none;font-family:FiraCode-Regular,monospace,sans-serif;font-weight:500;text-align:center;line-height:2.2em;margin:0 5px;transition:all .2s;width:64px;height:34px;background:rgba(0,0,0,0);border:1.8px solid #1c1c1c;color:#1c1c1c;font-size:14px;line-height:2.2em;width:auto}.tbtn:hover{color:#ccc;background-color:#1c1c1c}.tbtn a{color:#1c1c1c}.dropdown-content{background-color:#1c1c1c}.dropdown-content a{color:#fff}.dropdown-content a:hover{background-color:#333}.sidebar{background-color:#1c1c1c;color:#ccc}.sidebar a:hover{color:#f8efe1}body{background-color:#f8efe1;color:#1c1c1c}h1,h2,h3,h4,h5,h6{color:#000}code{background-color:#f8f8f8;border:1px solid #eaeaea}pre{background-color:#f8f8f8;border:1px solid #ccc}pre.highlight{background-color:#1c1c1c;border:1px solid #ccc}.token-del{background:#c00000;color:#f8f8f8}.token-add{background:green;color:#f8f8f8}`

)
