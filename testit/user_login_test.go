package testit_test

import (
	//. "path/to/potato"

	"github.com/goatcms/goatcms/cmsapp/dao"
	"github.com/goatcms/goatcms/cmsapp/entities"
	"github.com/goatcms/goatcms/cmsapp/services"
	"github.com/goatcms/goatcms/testit/itbase"
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/appdo"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var (
	testID       = int64(1)
	testUsername = "testusername"
	testPassword = "test-password"
)

var _ = Describe("UserLogin", func() {
	var (
		page *agouti.Page
		mapp app.App
	)

	BeforeEach(func() {
		var (
			err  error
			hash string
			deps struct {
				Crypt   services.Crypt   `dependency:"CryptService"`
				Fixture services.Fixture `dependency:"FixtureService"`
				Updater dao.UserUpdate   `dependency:"UserUpdate"`
			}
		)
		mapp, err = itbase.NewITApp(":3000")
		Expect(err).NotTo(HaveOccurred())
		err = mapp.DependencyProvider().InjectTo(&deps)
		Expect(err).NotTo(HaveOccurred())
		// load fixtures
		err = itbase.CreateSchema(mapp)
		Expect(err).NotTo(HaveOccurred())
		err = itbase.LoadFixtures(mapp, "random")
		Expect(err).NotTo(HaveOccurred())
		// set test credentials
		hash, err = deps.Crypt.Hash(testPassword)
		Expect(err).NotTo(HaveOccurred())
		err = deps.Updater.Update(mapp.AppScope(), &entities.User{
			ID:       &testID,
			Username: &testUsername,
			Password: &hash,
		}, &entities.UserFields{
			ID:       true,
			Username: true,
			Password: true,
		})
		Expect(err).NotTo(HaveOccurred())
		err = mapp.AppScope().Trigger(app.CommitEvent, nil)
		Expect(err).NotTo(HaveOccurred())
		// Run
		itbase.Run(mapp)
		// page prepare
		page, err = agoutiDriver.NewPage()
		Expect(err).NotTo(HaveOccurred())
		page.Size(1024, 768)
	})

	AfterEach(func() {
		err := mapp.AppScope().Trigger(app.CloseEvent, nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(page.Destroy()).To(Succeed())
		Expect(appdo.Close(mapp)).To(Succeed())
	})

	It("should manage user authentication", func() {

		By("go the user to the login form from the home page", func() {
			Expect(page.Navigate("http://localhost:3000")).To(Succeed())
			//fmt.Println(page.HTML())
			Expect(page.FindByID("signin-btn").Click()).To(Succeed())
		})

		By("allowing the user to fill out the login form and submit it", func() {
			Eventually(page.FindByName("Username")).Should(BeFound())
			Expect(page.FindByName("Username").Fill(testUsername)).To(Succeed())
			Expect(page.FindByName("Password").Fill(testPassword)).To(Succeed())
			Expect(page.Find("form.custom-signin-form").Submit()).To(Succeed())
		})
	})
})
