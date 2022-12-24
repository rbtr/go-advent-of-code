$(DIR)/main.go:
	cp skel/* $(DIR)/

$(DIR)/input:
	touch $(DIR)/sample $(DIR)/input
	 
.DEFAULT:
	mkdir -p $@
	make DIR=$@ $@/main.go $@/input 
