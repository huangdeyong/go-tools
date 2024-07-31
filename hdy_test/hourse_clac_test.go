package hdy

//
//import (
//	"fmt"
//	"math"
//	"testing"
//)
//
//func Test02_0(t *testing.T) {
//	// Get user input for loan amount, interest yearRate, and limitYear
//	var total float64 = 10000
//	var yearRate float64 = 3.5
//	var limitYear int = 5
//
//	//fmt.Print("Enter loan amount: ")
//	//fmt.Scanln(&total)
//	//
//	//fmt.Print("Enter interest yearRate (e.g., 5.5 for 5.5%): ")
//	//fmt.Scanln(&yearRate)
//	//
//	//fmt.Print("Enter loan limitYear in years: ")
//	//fmt.Scanln(&limitYear)
//
//	// Check for invalid loan limitYear (zero)
//	if limitYear <= 0 {
//		fmt.Println("Error: Loan limitYear must be greater than zero.")
//		return
//	}
//
//	// Calculate monthly interest yearRate
//	monthlyRate := yearRate / 100 / 12
//
//	// Use math.Pow for more accurate exponentiation (handles small interest rates)
//	monthlyPayment := total * monthlyRate * math.Pow(1+monthlyRate, float64(limitYear*12)) / (math.Pow(1+monthlyRate, float64(limitYear*12)) - 1)
//
//	// Display monthly payment
//	fmt.Printf("Monthly payment: %.2f\n", monthlyPayment)
//}
//
//// 计算房贷每月还款和总还款
//func Test02(t *testing.T) {
//	// 获取用户输入的贷款金额、年利率和贷款期限
//	var total float64 = 10000       // 贷款总额（默认10000元）
//	var yearRate float64 = 3.5      // 年利率（默认3.5%）
//	var limitYear int = 5           // 贷款期限（默认5年）
//	var prepayYear int = 2          // 提前还款年限（0表示不提前还款）
//	var prepayAmount float64 = 1000 // 提前还款金额
//
//	//// 提示用户输入可选的提前还款信息
//	//fmt.Println("请输入贷款总额（默认10000元）：")
//	//fmt.Scanln(&total)
//	//
//	//fmt.Println("请输入年利率（例如5.5表示5.5%，默认3.5%）：")
//	//fmt.Scanln(&yearRate)
//	//
//	//fmt.Println("请输入贷款期限（单位：年，默认5年）：")
//	//fmt.Scanln(&limitYear)
//	//
//	//fmt.Println("请输入提前还款年限（0表示不提前还款）：")
//	//fmt.Scanln(&prepayYear)
//	//
//	//if prepayYear > 0 {
//	//	fmt.Println("请输入提前还款金额：")
//	//	fmt.Scanln(&prepayAmount)
//	//}
//
//	// 检查贷款期限是否有效（大于零）
//	if limitYear <= 0 {
//		fmt.Println("错误：贷款期限必须大于零。")
//		return
//	}
//
//	// 计算每月利率
//	monthlyRate := yearRate / 100 / 12
//
//	// 计算不提前还款情况下的总还款
//	totalPaymentWithoutPrepay := calcTotalPayment(total, monthlyRate, limitYear)
//
//	// 处理提前还款（如果适用）
//	var totalPaymentWithPrepay float64
//	var interestSaved float64
//	if prepayYear > 0 && prepayAmount > 0 {
//		// 计算提前还款后剩余贷款金额
//		remainingLoan := total - prepayAmount - calcTotalBenJin(total, monthlyRate, limitYear*12)
//
//		// 计算剩余还款年限
//		remainingYears := limitYear - prepayYear
//
//		// 计算提前还款情况下的总还款
//		totalPaymentWithPrepay = calcTotalPayment(remainingLoan, monthlyRate, remainingYears)
//
//		// 计算节省的利息
//		interestSaved = totalPaymentWithoutPrepay - totalPaymentWithPrepay
//	} else {
//		totalPaymentWithPrepay = totalPaymentWithoutPrepay
//	}
//
//	// 显示结果
//	fmt.Printf("不提前还款情况下的每月还款：%.2f元\n", totalPaymentWithoutPrepay)
//	fmt.Printf("不提前还款情况下的总还款：%.2f元\n", totalPaymentWithoutPrepay*float64(limitYear*12))
//	fmt.Printf("提前还款情况下的每月还款：%.2f元\n", totalPaymentWithPrepay)
//	fmt.Printf("提前还款情况下的总还款：%.2f元\n", totalPaymentWithPrepay*float64((limitYear-prepayYear)*12)+prepayAmount)
//	fmt.Printf("提前还款节省的利息：%.2f元\n", interestSaved)
//}
//
//// 计算总还款
//func calcTotalPayment(loanAmount float64, monthlyRate float64, years int) float64 {
//	return loanAmount * monthlyRate * math.Pow(1+monthlyRate, float64(years*12)) / (math.Pow(1+monthlyRate, float64(years*12)) - 1)
//}
//
//func ()  {
//
//}
