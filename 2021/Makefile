$(DIR)/main.go:
	cp main.tmpl $(DIR)/main.go

$(DIR)/sample:
	touch $(DIR)/sample

$(DIR)/input:
	touch $(DIR)/input
	 
.DEFAULT:
	mkdir -p $@
	make DIR=$@ $@/main.go $@/sample $@/input

