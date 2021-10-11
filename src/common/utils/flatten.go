package utils

func FlattenStringStruts (acc *[]interface{},  i interface{}, isArray bool) (error) {

	if isArray {
		 switch  i.(type){
		 case [][]string:
			 for _, val := range i.([][]string) {
				 err := FlattenStringStruts(acc, val, true)
				 if err != nil {
					 return err
				 }
			 }
		 case []string:
			 for _, val := range i.([]string) {
				 err := FlattenStringStruts(acc, val, true)
				 if err != nil {
					 return err
				 }
			 }
		 case string:
			 *acc = append(*acc, i)
			 return nil
		 }
	}
	return nil

}