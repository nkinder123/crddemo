##### 总结：这是一个使用code-genertor生成的目录结构，我们使用hack中的update-codegen.sh脚本调用vendor包中的code-generator generate-groups.sh脚本进行生成generotor的目录

- 首先我们定义

  - pkg/apis/crddemo/v1中的doc.go、types.go
  - hack/boilerplate.go.txt、hack/tools.go、hack/update-codegen.sh、hack/verify-codegen.sh

  （此时需要注意，在update-codegen.sh执行code-genertor中的generate-groups脚本的时候仅在code-genertorv.30以下进行支持，v.30以及以后的code-genertor版本使用kube_codegen.sh，需要在update-codegen.sh中进行更改）

- 执行hack/update-codegen.sh就可以对generate目录进行生成了，生成的目录结构如下：
