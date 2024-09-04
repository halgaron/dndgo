package main
import ("fmt")

func main() {
  var basecr string
  fmt.Print("Type your Expected CR: ") 
  fmt.Scanf("%v", &basecr)
//   fmt.Println("Your listed cr is:", basecr)
  switch basecr {
	case "0": 
	var profbonus,ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr,coredamagecr = 2,13,1,6,3,0,1,13,0,1,1
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr,"Core Damage CR:",coredamagecr)
	case "1/8": 
	var profbonus,ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr,coredamagecr = 2,  13,  7,  35,  3, 2,  3,  13,  0.125,  2, 2
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "1/4": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 2,  13, 36,  49,  3, 4,  5,  13,  0.25, 3, 3 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "1/2": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 2,  13,  50,  70,  3,   6,  8,  13,  0.5,  4, 4 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "1": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 2,  13,    71,  85,  3,  9,  14,  13, 1, 5, 5
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "2": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 2,  13,  86,  100,  3,   15,  20,  13,  2,  6, 6
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "3": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 2,  13,   101,  115,  4,  21,  26,  13,  3,  7, 7
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "4": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 2,   14,    116,  130,  5,   27,  32,  14,  4,  8, 8
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "5": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 3,  15,   131,  145,  6,   33,  38,  15,  5,  9, 9
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "6":
    var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 3,   15,    146,  160,  6,   39,  44,  15,  6,  10, 10
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "7": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 3,   15,    161,  175,  6,    45,  50,  15,  7,  11, 11
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "8": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 3,   16,    176,  190,  7,    51,  56,  16,  8,  12, 12
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "9": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 4,   16,   191,  205,  7,    57,  62,  16,  9,  13, 13
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "10": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 4,   17,  206,  220,  7,   63,  68,  16,  10,  14, 14
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "11": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 4,   17,  221,  235,  8, 69,  74,  17,  11,  15, 15
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "12": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 4,   17,   236,  250,  8,   75,  80,  17,  12,  16, 16
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "13": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 5,   18,   251,  265,  8,   81,  86,  18,  13,  17, 17
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "14": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 5,   18,    266,  280,  8,  87,  92,  18,  14,  18, 18
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "15": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 5,   18,   281,  295,  8,  93,  98,  18,  15,  19, 19
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "16": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 5,   18,   295,  310,  9,    99,  104,  18,  16,  20, 20 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "17": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 6,   19,   311,  325,  10,  105,  110,  19,  17,  21, 21 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "18": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 6,   19,  326,  340,  10,   111,  116,  19,  18,  22, 22 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "19": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 6,   19,    341,  355,  10, 117,  122,  19,  19,  23, 23 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "20": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 6,   19,  356,  400,  10,  123,  140,  19,  20,  24, 24 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "21": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 7,   19,   401,  445,  11,   141,  158,  20,  21,  25, 25
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "22": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 7,   19,   446,  490,  11,  159,  176,  20,  22,  26, 26
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "23": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 7,   19,  491,  535,  11, 177,  194,  20,  23,  27, 27 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "24": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 7,   19,   536,  580,  12,    195,  212,  21,  24,  28, 28 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "25": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 8,   19,   581,  625,  12,   213,  230,  21,  25,  29, 29
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "26": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 8,   19,   626,  670,  12,  231,  248,  21,  26,  30, 30 
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "27": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 8,   19,   671,  715,  13,    249,  266,  22,  27,  31, 31
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "28": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 8,   19,    716,  760,  13,  267,  284,  22,  28,  32, 32
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "29": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 9,   19,   761,  805,  13,   285,  302,  22,  29,  33, 33
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	case "30": 
	var profbonus, ac,hpmin,hpmax,attackbonus,damagemin,damagemax,savedc,corecr,corehpcr, coredamagecr = 9,   19,  806,  850,  14,   303,  320,  23,  30,  34, 34
	fmt.Println("AC:", ac, "HP", hpmin,"-",hpmax)
	fmt.Println("AttackBonus:", attackbonus, "Prof Bonus:+",profbonus,"Damage", damagemin,"-",damagemax)
	fmt.Println("SaveDC:", savedc)
	fmt.Println("CoreCR:", corecr, "corehpcr", corehpcr, "Coredamagecr:",coredamagecr)
	}

}

func dothing(){
  var baseac int
  var basehp int
  fmt.Print("Type your AC:","Type your HP:") 
  fmt.Scanf("%v","%v", &baseac,&basehp)
}

// func source(ac *int) {
// var blah = (&ac)

// }

// func grab() {
	
// 	fmt.Println("whatisblah",source(*blah))
// }