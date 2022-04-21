package main

import (
	"github.com/Knetic/govaluate"
	"github.com/spf13/viper"
	"log"
)

func init() {
	// godotenv.Load(".env")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
func main() {
	//router := gin.Default()
	//router.POST("/reversearray", func(c *gin.Context) {
	//	inputInts := []int{}
	//	c.ShouldBind(&inputInts)
	//
	//})
	//
	//router.Run(":" + viper.Get("PORT").(string))
	expression, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		log.Println(err.Error())
	}
	result, err := expression.Evaluate(nil)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(result)
}
