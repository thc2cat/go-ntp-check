TAG=$(shell git tag)
NAME= $(notdir $(shell pwd))
SOURCES= $(wildcard *.go)
OPTIONS=-v -server ntp.uvsq.fr

CLEANLIST=${NAME} ${NAME}-${TAG} ${NAME}.exe

all:	${NAME}  

${NAME}:  ${SOURCES}
	go build -ldflags '-w -s -X main.Version=NTPCheck-${TAG}' 

clean:
	@touch ${CLEANLIST}
	@rm ${CLEANLIST}
	

test:  ${NAME} Makefile
	 ./${NAME} ${OPTIONS} 

upx:	${NAME}
	goupx ${NAME}
