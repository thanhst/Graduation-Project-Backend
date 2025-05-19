package helper

import (
	"math/rand"
	"server/internal/utils/dotenv"
)

var names = []string{
	"Nguyen Van An", "Tran Thi Bich", "Le Hoang Nam", "Pham Minh Chau", "Hoang Gia Bao",
	"Dang Thi Mai", "Vo Van Khoa", "Bui Thi Lan", "Do Anh Tuan", "Ngo Thi Hoa",
	"Nguyen Thi Huong", "Tran Van Hieu", "Le Thi Kim", "Pham Van Duc", "Hoang Minh Tien",
	"John Smith", "Emily Johnson", "Michael Brown", "Jessica Williams", "David Miller",
	"Sarah Davis", "James Wilson", "Ashley Moore", "Robert Taylor", "Olivia Anderson",
	"William Thomas", "Emma Jackson", "Joseph White", "Sophia Harris", "Daniel Martin",
}

var emails = []string{
	"john.smith@example.com",
	"mary.jones@example.com",
	"robert.brown@example.com",
	"linda.nguyen@example.com",
	"michael.tran@example.com",
	"james.le@example.com",
	"patricia.pham@example.com",
	"jennifer.ho@example.com",
	"william.dang@example.com",
	"david.pham@example.com",
	"richard.nguyen@example.com",
	"joseph.le@example.com",
	"thomas.nguyen@example.com",
	"charles.pham@example.com",
	"christopher.tran@example.com",
	"daniel.ho@example.com",
	"matthew.dang@example.com",
	"anthony.le@example.com",
	"mark.nguyen@example.com",
	"donald.pham@example.com",
	"steven.tran@example.com",
	"paul.ho@example.com",
	"andrew.dang@example.com",
	"joshua.le@example.com",
	"kevin.nguyen@example.com",
	"brian.pham@example.com",
	"george.tran@example.com",
	"edward.ho@example.com",
	"ronald.dang@example.com",
	"timothy.le@example.com",
	"jason.nguyen@example.com",
	"jeffrey.pham@example.com",
	"ryan.tran@example.com",
	"jacob.ho@example.com",
	"gary.dang@example.com",
	"nicholas.le@example.com",
	"eric.nguyen@example.com",
	"stephen.pham@example.com",
	"jonathan.tran@example.com",
	"larry.ho@example.com",
	"justin.dang@example.com",
	"scott.le@example.com",
	"brandon.nguyen@example.com",
	"benjamin.pham@example.com",
	"samuel.tran@example.com",
	"gregory.ho@example.com",
	"frank.dang@example.com",
	"alexander.le@example.com",
	"raymond.nguyen@example.com",
	"patrick.pham@example.com",
	"jack.tran@example.com",
	"dennis.ho@example.com",
	"jerry.dang@example.com",
	"tyler.le@example.com",
	"aaron.nguyen@example.com",
	"joe.pham@example.com",
	"philip.tran@example.com",
	"willie.ho@example.com",
	"billy.dang@example.com",
	"bryan.le@example.com",
	"lois.nguyen@example.com",
	"teresa.pham@example.com",
	"nancy.tran@example.com",
	"karen.ho@example.com",
	"helen.dang@example.com",
	"donna.le@example.com",
	"carol.nguyen@example.com",
	"ruth.pham@example.com",
	"sharon.tran@example.com",
	"michelle.ho@example.com",
	"laura.dang@example.com",
	"sarah.le@example.com",
	"kimberly.nguyen@example.com",
	"deborah.pham@example.com",
	"jessica.tran@example.com",
	"stephanie.ho@example.com",
	"emily.dang@example.com",
	"madison.le@example.com",
	"amber.nguyen@example.com",
	"brittany.pham@example.com",
	"morgan.tran@example.com",
	"megan.ho@example.com",
	"ashley.dang@example.com",
	"kathleen.le@example.com",
	"marie.nguyen@example.com",
	"julia.pham@example.com",
	"olivia.tran@example.com",
	"grace.ho@example.com",
	"eva.dang@example.com",
	"vivian.le@example.com",
	"daisy.nguyen@example.com",
	"han.nguyen@example.com",
	"quang.le@example.com",
	"hoa.pham@example.com",
	"thanh.tran@example.com",
	"huong.ho@example.com",
	"tuan.dang@example.com",
}

var descriptions = []string{
	"Software engineer with 5 years experience",
	"Specializes in backend development",
	"Enjoys hiking on weekends",
	"Graphic designer focused on minimalistic design",
	"Loves photography and traveling",
	"Coffee enthusiast",
	"Marketing specialist passionate about digital campaigns",
	"Project manager skilled in agile methodologies",
	"Data scientist focused on machine learning and AI",
	"Full-stack developer with React and Node.js experience",
}

var states = []string{
	"online",
	"offline",
}

var roles = []string{
	"teacher",
	"student",
}

var loginMethods = []string{
	"email",
	"google",
	"github",
}

var imageUsers = []string{
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/1.jpg",
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/2.jpg",
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/3.jpg",
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/4.jpg",
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/5.jpg",
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/6.jpg",
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/7.jpg",
	dotenv.GetDotEnv("APP_URL") + ":" + dotenv.GetDotEnv("APP_PORT") + "/uploads/8.jpg",
}

func RandomName() string {
	return names[rand.Intn(len(names))]
}

func RandomEmail() string {
	return emails[rand.Intn(len(emails))]
}
func RandomDescription() string {
	return descriptions[rand.Intn(len(descriptions))]
}
func RandomState() string {
	return states[rand.Intn(len(states))]
}
func RandomRole() string {
	return roles[rand.Intn(len(roles))]
}

func RandomLoginMethod() string {
	return loginMethods[rand.Intn(len(loginMethods))]
}

func RandomeImagesURL() string {
	return imageUsers[rand.Intn(len(imageUsers))]
}
