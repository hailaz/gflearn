# 清空文件
> README.md
# 写入基础内容
cat base.md | while read line; do echo $line >> README.md; done
# 写入例子目录
for file in $(find . -name go.mod); do
    # echo $file
    dirpath=$(dirname $file)
    title=$(head -n +1 ${dirpath}/readme.md)
    # echo ${title:2}
    echo -e "- [${title:2}](${dirpath:2})\n" >> README.md
done
