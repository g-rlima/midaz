package portfolio

import (
	"bytes"
	"testing"
	"time"

	"github.com/LerianStudio/midaz/common/mmodel"
	"github.com/LerianStudio/midaz/components/mdz/internal/domain/repository"
	"github.com/LerianStudio/midaz/components/mdz/pkg/factory"
	"github.com/LerianStudio/midaz/components/mdz/pkg/iostreams"
	"github.com/LerianStudio/midaz/components/mdz/pkg/ptr"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gotest.tools/golden"
)

func Test_newCmdPortfolioDescribe(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := repository.NewMockPortfolio(ctrl)

	ledgerID := "0192e251-328d-7390-99f5-5c54980115ed"
	organizationID := "0192e250-ed9d-7e5c-a614-9b294151b572"
	portfolioID := "01931b44-6e33-791a-bfad-27992fa15984"

	ledFactory := factoryPortfolioDescribe{
		factory: &factory.Factory{IOStreams: &iostreams.IOStreams{
			Out: &bytes.Buffer{},
			Err: &bytes.Buffer{},
		}},
		repoPortfolio:  mockRepo,
		OrganizationID: organizationID,
		LedgerID:       ledgerID,
		Out:            "",
		JSON:           false,
	}

	metadata := map[string]any{
		"chave1": "valor1",
		"chave2": 2,
		"chave3": true,
	}

	cmd := newCmdPortfolioDescribe(&ledFactory)
	cmd.SetArgs([]string{
		"--ledger-id", ledgerID,
		"--organization-id", organizationID,
		"--portfolio-id", portfolioID,
	})

	item := mmodel.Portfolio{
		ID:             portfolioID,
		Name:           "2Real",
		OrganizationID: organizationID,
		LedgerID:       ledgerID,
		Status: mmodel.Status{
			Code:        "ACTIVE",
			Description: ptr.StringPtr("Teste Ledger"),
		},
		CreatedAt: time.Date(2024, 10, 31, 11, 23, 45, 165229000, time.UTC),
		UpdatedAt: time.Date(2024, 10, 31, 11, 23, 45, 165229000, time.UTC),
		DeletedAt: nil,
		Metadata:  metadata,
	}

	mockRepo.EXPECT().GetByID(gomock.Any(), gomock.Any(), gomock.Any()).Return(&item, nil)

	err := cmd.Execute()
	assert.NoError(t, err)

	output := ledFactory.factory.IOStreams.Out.(*bytes.Buffer).Bytes()
	golden.AssertBytes(t, output, "output_describe.golden")
}