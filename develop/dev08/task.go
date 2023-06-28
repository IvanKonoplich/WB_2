package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mitchellh/go-ps"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
//
//
//- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
//- pwd - показать путь до текущего каталога
//- echo <args> - вывод аргумента в STDOUT
//- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
//- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
//
//
//
//
//Так же требуется поддерживать функционал fork/exec-команд
//
//
//Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
//
//
//*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
//в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
//и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

func main() {
	shell()
}

func shell() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		//определяем команду
		scanner.Scan()
		cmd := scanner.Text()
		if strings.Contains(cmd, "|") { //если содержит |, то является конвейером
			if err := pipes(cmd); err != nil {
				fmt.Println(err.Error())
			}
		} else {
			if err := execCmd(cmd); err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func execCmd(cmd string) error {
	c := strings.Split(cmd, " ")
	switch c[0] {
	case "cd":
		if len(c) < 2 {
			var err error
			c[1], err = os.UserHomeDir()
			if err != nil {
				return err
			}
		}
		return os.Chdir(c[1])
	case "pwd":
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(path)
	case "echo":
		for i := 1; i < len(c); i++ {
			fmt.Print(c[i], " ")
		}
	case "kill":
		if len(c) < 2 {
			return fmt.Errorf("недостаточно аргументов для kill")
		}

		pid, err := strconv.Atoi(c[1])
		if err != nil {
			p, err := ps.Processes()
			if err != nil {
				return err
			}
			for _, j := range p {
				if j.Executable() == c[1] {
					pid = j.Pid()
					break
				}
			}
			if pid == 0 {
				return fmt.Errorf("процесс: '%s' не существует", c[1])
			}
		}
		p, err := os.FindProcess(pid)
		if err != nil {
			return err
		}

		return p.Kill()
	case "ps":
		p, err := ps.Processes()
		if err != nil {
			return err
		}
		for _, j := range p {
			fmt.Println(j.Pid(), j.Executable())
		}
	case "fork":
		go execCmd(strings.Join(c[1:], " "))
		execCmd(strings.Join(c[1:], " "))
	case "exec":
		go execCmd(strings.Join(c[1:], " "))
	case "quit":
		os.Exit(0)
	}
	return nil
}

func pipes(cmd string) error {
	c := strings.Split(cmd, " | ") //определяем команды в конвейере
	if len(c) < 2 {
		return fmt.Errorf("недостаточное колличество аргументов в конвеере: '%v'", c)
	} //если только dev01 команда, то возвращаем ошибку

	var b bytes.Buffer
	for i := 0; i < len(c); i++ { //пробегаемся по командам
		com := exec.Command(c[i])              //подготавливаем команду
		com.Stdin = bytes.NewReader(b.Bytes()) //входные данные команды берутся из буфера в который пишет предыдущая команда
		b.Reset()
		com.Stdout = &b //пишем в буфер

		err := com.Run() //выполняем команду
		if err != nil {
			return err
		}
	}

	fmt.Fprint(os.Stdout, b.String()) //выводим результат
	return nil
}
