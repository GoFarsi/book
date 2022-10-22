.PHONY: pdf epub html

pdf:
	mkdir -p build
	pandoc \
    --pdf-engine=xelatex \
    --template=./assets/templates/eisvogel.latex \
    --highlight-style tango \
    --toc -N \
    --filter pandoc-crossref \
    --metadata rtl=true \
    -o build/ouput.pdf \
    ./info.txt src/*.md