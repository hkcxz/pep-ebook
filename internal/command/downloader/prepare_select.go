package downloader

import (
	"errors"
	"strings"

	"github.com/maogou/pep-ebook/internal/classification"

	"github.com/AlecAivazis/survey/v2"
)

func (d *Downloader) prepareSelect() error {

	var (
		period  string
		grade   string
		subject string
	)

	prompt := &survey.Select{
		Message: "请选择用书学段",
		Options: classification.Periods,
	}

	if err := survey.AskOne(prompt, &period); err != nil {
		return errors.New("你中断了选择用书学段")
	}

	prompt = &survey.Select{
		Message: "请选择学生年级",
		Options: classification.Grades[period],
	}

	if err := survey.AskOne(prompt, &grade); err != nil {
		return errors.New("你中断了选择学生年级")
	}

	prompt = &survey.Select{
		Message: "请选择学科",
		Options: classification.Subjects[period],
	}

	if err := survey.AskOne(prompt, &subject); err != nil {
		return errors.New("你中断了选择学科")
	}

	key := period + "-" + grade + "-" + subject

	if _, ok := classification.Paths[key]; !ok {
		return errors.New("你选择的学段+年级+学科不存在,请重新选择")
	}

	if len(classification.Paths[key]) == 0 {
		return errors.New("你选择的学段+年级+学科不存在,请重新选择")
	}

	d.period = period
	d.grade = grade
	d.subject = subject
	d.paths = classification.Paths[key]
	d.pathKey = strings.ReplaceAll(key, "-", "/")

	return nil
}
