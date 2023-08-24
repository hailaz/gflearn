"use strict";(self.webpackChunkhailaz_github_io=self.webpackChunkhailaz_github_io||[]).push([[100],{5728:(t,e,a)=>{a.r(e),a.d(e,{assets:()=>c,contentTitle:()=>s,default:()=>m,frontMatter:()=>o,metadata:()=>r,toc:()=>d});var n=a(7462),i=(a(7294),a(3905)),l=a(4915);const o={},s="\u9759\u6001\u6587\u4ef6",r={unversionedId:"web/static/readme",id:"web/static/readme",title:"\u9759\u6001\u6587\u4ef6",description:"\u4ee3\u7801 - main.go",source:"@site/docs/web/1-static/readme.md",sourceDirName:"web/1-static",slug:"/web/static/",permalink:"/gflearn/docs/web/static/",draft:!1,editUrl:"https://github.com/hailaz/gflearn/blob/main/docs/web/1-static/readme.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"\u4f60\u597d\uff0c\u4e16\u754c",permalink:"/gflearn/docs/web/helloworld/"},next:{title:"\u4e0a\u4f20\u6587\u4ef6",permalink:"/gflearn/docs/web/uploadfile/"}},c={},d=[{value:"\u4ee3\u7801 - main.go",id:"\u4ee3\u7801---maingo",level:2},{value:"\u914d\u7f6e - config.yaml",id:"\u914d\u7f6e---configyaml",level:2},{value:"\u9759\u6001\u6587\u4ef6 - ./static/1.txt",id:"\u9759\u6001\u6587\u4ef6---static1txt",level:2},{value:"\u542f\u52a8",id:"\u542f\u52a8",level:2},{value:"\u8bbf\u95ee",id:"\u8bbf\u95ee",level:2}],g={toc:d};function m(t){let{components:e,...a}=t;return(0,i.kt)("wrapper",(0,n.Z)({},g,a,{components:e,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"\u9759\u6001\u6587\u4ef6"},"\u9759\u6001\u6587\u4ef6"),(0,i.kt)("h2",{id:"\u4ee3\u7801---maingo"},"\u4ee3\u7801 - main.go"),(0,i.kt)(l.Z,{language:"go",mdxType:"CodeBlock"},'package main\n\nimport (\n\t"github.com/gogf/gf/v2/frame/g"\n\t"github.com/gogf/gf/v2/net/ghttp"\n)\n\nfunc main() {\n\ts := g.Server()\n\ts.AddStaticPath("/static", "static")\n\ts.BindHandler("/", func(r *ghttp.Request) {\n\t\tr.Response.Write("helloworld")\n\t})\n\t// s.SetAddr("127.0.0.1:8080") // \u770bconfig.toml\u914d\u7f6e\u6587\u4ef6\n\ts.Run()\n}\n\n// http://127.0.0.1:8080/static/1.txt\n'),(0,i.kt)("h2",{id:"\u914d\u7f6e---configyaml"},"\u914d\u7f6e - config.yaml"),(0,i.kt)(l.Z,{language:"yaml",mdxType:"CodeBlock"},'server:\n    address: "127.0.0.1:8080"\n'),(0,i.kt)("h2",{id:"\u9759\u6001\u6587\u4ef6---static1txt"},"\u9759\u6001\u6587\u4ef6 - ./static/1.txt"),(0,i.kt)(l.Z,{language:"text",mdxType:"CodeBlock"},"\u6211\u662f\u9759\u6001\u6587\u4ef6\u5185\u5bb9"),(0,i.kt)("h2",{id:"\u542f\u52a8"},"\u542f\u52a8"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-bash"},"go run main.go\n")),(0,i.kt)("h2",{id:"\u8bbf\u95ee"},"\u8bbf\u95ee"),(0,i.kt)("p",null,(0,i.kt)("a",{parentName:"p",href:"http://127.0.0.1:8080/static/1.txt"},"http://127.0.0.1:8080/static/1.txt")))}m.isMDXComponent=!0}}]);