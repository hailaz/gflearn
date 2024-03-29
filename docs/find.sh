#!/bin/bash
rmname="readme.md"
menu="0-readme.md"
# 清空文件
> ${menu}
# 写入基础内容
cat base.txt | while read line; do echo $line >> ${menu}; done
# 写入例子目录
for subdir in $(ls .); do
    if [ -d ${subdir} ]; then
        # 写入目录
        title=$(head -n +1 ${subdir}/${rmname})
        echo -e "- [${title:2}](${subdir})" >> ${menu}
        echo "====find subdir: ${subdir}"
        # 写入项目
        for file in $(find ${subdir} -name ${rmname}); do
            if [ "${subdir}/${rmname}" != "${file}" ]; then
                projectpath=$(dirname ${file})
                title=$(head -n +1 ${projectpath}/${rmname})
                echo "====find file: ${file}"
                echo "  - [${title:2}](${subdir}/${projectpath#*-})"
                echo -e "  - [${title:2}](${subdir}/${projectpath#*-})" >> ${menu}
            fi
        done
    fi
done
