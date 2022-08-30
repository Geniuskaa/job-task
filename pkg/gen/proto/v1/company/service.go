package company

import (
	"context"
	"fmt"
	company "github.com/Geniuskaa/job-task/pkg/gen/proto/v1"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

type Service struct {
	company.UnimplementedRusProfileServer
}

func NewService() *Service {
	return &Service{}
}

func (s Service) GetCompanyInfo(ctx context.Context, in *company.CompanyINN) (*company.CompanyInfo, error) {

	comp, err := rusProfileCompanyInfo(in.Inn)
	if err != nil {
		return nil, err
	}

	return comp, nil
}

func rusProfileCompanyInfo(inn string) (*company.CompanyInfo, error) {
	client := http.Client{}

	res, err := client.Get(fmt.Sprintf("https://www.rusprofile.ru/search?query=%s", inn))
	if err != nil {
		log.Println(fmt.Errorf("rusProfileCompanyInfo failed: %w", err))
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(fmt.Errorf("rusProfileCompanyInfo failed: %w", err))
		return nil, err
	}

	kpp := doc.Find("#clip_kpp").Text()
	compName := doc.Find(".company-header__row h1").Text()
	compName = strings.ReplaceAll(compName, "\n", "")
	compName = strings.ReplaceAll(compName, "\"", "")
	compName = strings.TrimSpace(compName)
	ownerName := doc.Find(".gtm_main_fl span").Text()

	comp := company.CompanyInfo{
		Inn:           inn,
		Kpp:           kpp,
		CompanyName:   compName,
		OwnerFullName: ownerName,
	}

	return &comp, nil
}
