# 清空文件
>README.md
# 写入基础内容
cat base.md | while read line; do echo $line >>README.md; done
# 写入例子目录
for subdir in $(ls .); do
    if [ -d ${subdir} ]; then
        # echo ${subdir}
        title=$(head -n +1 ${subdir}/readme.md)
        echo -e "- [${title:2}](${subdir})" >>README.md
        for file in $(find ${subdir} -name go.mod); do
            # echo $file
            projectpath=$(dirname $file)
            title=$(head -n +1 ${projectpath}/readme.md)
            # echo ${title:2}
            echo -e "   - [${title:2}](${projectpath})" >>README.md
        done
    fi
done
