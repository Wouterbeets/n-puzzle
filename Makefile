# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: wbeets <wbeets@student.42.fr>              +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2013/12/22 21:27:45 by wbeets            #+#    #+#              #
#    Updated: 2013/12/22 21:27:47 by wbeets           ###   ########.fr        #
#                                                                              #
# **************************************************************************** #
#
#NAME	= fdf
#SRCS	= main.c\
#		  get_data.c\
#		  calc.c\
#		  draw.c
#OBJS	= ${SRCS:.c=.o}
#INC		= ./
#FLAGS	= -Wall
#

all: 
	go clean &&	go build && ./gen 3 -u > boardExample.txt && cat boardExample.txt && ./n-puzzle -f boardExample.txt 

test:
	./gen 3 -u > unsolv3.txt
	./gen 3 -s > solv3.txt
	./gen 4 -u > unsolv4.txt
	./gen 4 -u > solv4.txt
	echo "testing size 3 solvable"
	go clean && go build && ./n-puzzle -f solv3.txt
	echo "testing size 3 unsolvable"
	go clean && go build && ./n-puzzle -f unsolv3.txt
	echo "testing size 4 solvable"
	go clean && go build && ./n-puzzle -f solv4.txt
	echo "testing size 4 unsolvable"
	go clean && go build && ./n-puzzle -f unsolv4.txt
