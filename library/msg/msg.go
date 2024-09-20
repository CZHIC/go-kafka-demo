package msg

var (
	CategoryMap map[int]string
	//文件夹名称与投后项目ID，对应关系
	NameToProjectId_Test    map[string]int
	NameToProjectId_Release map[string]int
	NameToProjectId_Pro     map[string]int
	ResolutionCategory      map[int]string
	DirectorMeeting         map[int]string
	TencentOpinion          map[int]string
	Conclusion              map[int]string
	UrgencyOrImportance     map[int]string
	FileTag                 map[int]string

	TeamUserParamMap  map[string]string
	SituationParamMap map[string]string
)

const (
	//project_flow_user
	InvestProject            = 1  // 投资项目
	CollaborationProject     = 2  // 合作项目
	InvestmentProject        = 3  // 投后项目
	CompanyProject           = 4  // 公司
	InvestmentProjectHistory = 5  // 投后历史项目
	InvestmentReview         = 10 // 投后review

	SortPositive = 1 //正序
	SortReverse  = 2 //倒序
	SortDefault  = 0 //默认

	//变更历史
	InvestmentInfo       = 31 // 投后基本信息
	ProductInfo          = 32 // 产品信息
	InvestmentTeamMember = 38 // 投后团队成员
	SituationInfo        = 42 // 产品管线信息

	// 盘点状态枚举
	InventoryInitial = 1 //盘点初始
	InventoryStop    = 2 //盘点中止
	InventoryDo      = 3 //盘点进行中
	InventoryAdvent  = 4 //盘点临期
	InventoryFinish  = 5 //盘点完成
	InventoryUrge    = 6 //催办

	// 权限状态码
	FunctionStatusValid       = 1 // 可见 / 不受限
	FunctionStatusDeactivate  = 2 // 可见不点
	FunctionStatusZero        = 3 // 无权限
	FunctionStatusStakeholder = 4 // 干系人受限
	FunctionStatusRegion      = 5 // 区域受限

	// t_stakeholder_source
	StakeholderSourceUserTpye  = 1 // 干系人
	StakeholderSourceTeamType  = 2 // 团队
	StakeholderSourceWhiteType = 3 // 白名单
)

func init() {
	CategoryMap = map[int]string{
		1: "Memo",
		2: "Model",
		3: "TS",
		4: "LF",
		5: "审批记录",
		6: "公司提供的资料",
		7: "DD", // 合并LDD 和 FDD  (尽职调查)
		//8:  "LDD",
		9:  "品类及竞品参考资料",
		99: "产品数据",
		10: "其他",
		11: "评测运营反馈",
		12: "会议纪要",
		13: "投后binder",
	}
	DirectorMeeting = map[int]string{
		1: "Board Resolution only",
		2: "Shareholder Resolution only",
		3: "Board Meeting&Resolution",
		4: "General Meeting&Resolution",
		5: "Company Update",
		6: "Other",
	}

	ResolutionCategory = map[int]string{
		1:  "Capital Operation (Financing/IPO/Liquidation/Spin-Off/Restructure)",
		2:  "M&A/Investments",
		3:  "Remuneration/IncentivePlan",
		4:  "Appointment/Change of Key Roles",
		5:  "Inter-Company Transactions",
		6:  "Budget/Financials",
		7:  "Strategy/Business Plan",
		8:  "License Application",
		9:  "Compliance/Lawsuit",
		10: "Other Major Management",
		11: "Other Administrative Support",
	}
	TencentOpinion = map[int]string{
		1: "Agree",
		2: "Disagree",
		3: "Reserved Opinion",
		4: "Disclaimer",
	}

	Conclusion = map[int]string{
		1: "Consent",
		2: "Rejection",
	}
	UrgencyOrImportance = map[int]string{
		1: "High",
		2: "Middle",
		3: "Low",
	}
	FileTag = map[int]string{
		1: "Agenda",
		2: "BOD Deck",
		3: "Meeting Minutes",
		4: "Board Resolution",
		5: "Agreement/Certificate",
		6: "Other",
		7: "General Meeting Deck",
		8: "Shareholder Resolution",
		9: "Power of Attorney",
	}
	TeamUserParamMap = map[string]string{
		"name":           "姓名",
		"job":            "职位",
		"tag":            "职级",
		"status":         "在职状态",
		"linkedIn":       "Linked In",
		"workExperience": "工作经历",
		"site":           "地址",
	}
	SituationParamMap = map[string]string{
		"productName":        "产品名称",
		"category":           "品类",
		"subProducts":        "子品类",
		"playAmplifier":      "玩法放大器",
		"culturalAmplifier":  "文化放大器",
		"platform":           "平台",
		"commercialModel":    "商业化模式",
		"ip":                 "IP",
		"ipAscription":       "IP归属",
		"publisher":          "发行商",
		"tencentCooperation": "与腾讯合作",
		"currentStage":       "当前阶段",
		"manpower":           "研发人力 (正编)",
		"manpowerOut":        "研发人力 (外包)",
		"costs":              "研发成本(百万)",
		"productRisk":        "产品风险",
		"riskDescription":    "风险说明",
		//"cooperationType":        "合作类型",
		//"periodCooperationBegin": "合作期限开始",
		//"periodCooperationEnd":   "合作期限结束",
		//"investCurrency":         "币种",
		//"contributionAmount":     "腾讯出资金额(百万)",
		//"lf":                     "LF (百万)",
		//"mg":                     "MG (百万)",
		"cooperationArea": "合作区域",
		//"dockingTeam":            "对接团队",
		//"directorName":           "负责人",
		"remark":                   "备注",
		"productExpectationList":   "产品预期",
		"goOnlineList":             "预期上线",
		"productCollaborationList": "合作信息",
		"productProgressList":      "进度信息",
		"competingProduct":         "对标竞品",
		"productDescription":       "产品描述",

		"engine":            "引擎",
		"artStyle":          "美术风格",
		"devPeriod":         "预计研发时长（月）",
		"banhaoAquired":     "已有版号",
		"abroadOpportunity": "出海机会",
		// 230320
		"title":            "产品简称",
		"code":             "产品代号",
		"otherName":        "产品其他名称",
		"modes":            "多人模式",
		"developer":        "研发商",
		"manufacturerType": "厂商类型",
		"ipOwner":          "ip所有者",
		"GAPPBanhao":       "中国版号",
		"engineType":       "引擎类型",
		"similarProduct":   "相似产品",
		"marketingBudget":  "市场预算",
		"totalBudget":      "总预算",
		"estYr1Rev":        "预估首年收入",
		"estLifetimeUnits": "预估销量",
		"tier1":            "tier1",
		"tier2":            "tier2",
		"tier3":            "tier3",
		"keySummary":       "要点总结",
		"keySummaryWriter": "要点总结-填写人",
		"keySummaryDate":   "要点总结-填写时间",
		"type":             "上线类型",
		"channel":          "上线渠道",
		// special
		"onlineArea":       "上线地区",
		"postpone":         "是否延期",
		"TBC":              "TBC",
		"expect":           "产品预期",
		"conclusion":       "是否为结论",
		"cooperationType":  "合作类型",
		"investCurrency":   "币种",
		"majorProduct":     "重点产品（portfolio层面）",
		"coreProduct":      "核心单品（厂商层面）",
		"goOnlinePlatform": "上线平台",
		// progress
		"meetingTime":             "日期",
		"stage":                   "阶段",
		"status":                  "状态",
		"fBD":                     "FBD",
		"fBDEvaluation":           "评估",
		"fBDDate":                 "评估",
		"ventureLab":              "Venture Lab",
		"ventureLabEvaluation":    "评估",
		"ventureLabDate":          "评估",
		"sponsorTeam":             "业务对接团队",
		"sponsorTeamEvaluation":   "评估",
		"sponsorTeamDate":         "评估",
		"eval":                    "评测团队",
		"evalEvaluation":          "评估",
		"evalDate":                "评估",
		"supportNeeded":           "所需支持",
		"supportNeededEvaluation": "评估",
		"supportNeededDate":       "评估",
	}

	NameToProjectId_Test = map[string]int{
		"大连测试公司":       113,
		"gabe-test001": 118,
		"GemiTest1221": 134,
		"文件测试":         1373,
		"转投后看文件":       1372,
	}

	NameToProjectId_Release = map[string]int{
		"预发布测试":      705,
		"czc-test":   706,
		"Angeltest":  707,
		"gemilitest": 708,
		"SSSSS":      710,
	}

	NameToProjectId_Pro = map[string]int{
		"Epic":                  384,
		"Robot":                 342,
		"Pocket Gems":           199,
		"Hi Rez":                372,
		"Discord":               389,
		"Bad Robot":             398,
		"Apponboard":            473,
		"WONDERSTORM":           318,
		"Wolfeye":               319,
		"Terrible Posture":      327,
		"Stoic":                 332,
		"Sphero":                334,
		"Red Hook":              345,
		"Otherside":             349,
		"offworld":              143,
		"Mohawk":                357,
		"Hellbent":              373,
		"Bunch Live":            376,
		"Gallium":               377,
		"Fenix Fire":            382,
		"Ember lab":             385,
		"Dreamcraft":            386,
		"Clapfoot":              392,
		"CAPYBARA":              394,
		"Brace Yourself":        395,
		"Shiny Shoe":            474,
		"QI Software":           475,
		"Airship":               490,
		"Raccoon":               467,
		"Eleventh hour":         512,
		"Heart Machine":         536,
		"Voodoo":                321,
		"Lockwood":              362,
		"Larian":                366,
		"BOHEMIA":               397,
		"4A Ukraine":            402,
		"Shiro":                 336,
		"Payload":               347,
		"Mundfish":              356,
		"Milky Tea":             359,
		"Critical Force":        391,
		"Antstream":             406,
		"22Cans":                141,
		"1939 Games":            403,
		"Triternion":            468,
		"Fall Damage":           471,
		"HAWKSWELL":             470,
		"False Prophet":         516,
		"Alter Games":           421,
		"Keen":                  166,
		"Exit plan":             179,
		"Playtonic":             416,
		"Vostok":                431,
		"Gaijin":                515,
		"Twin Earth":            450,
		"GFA":                   445,
		"Parasight":             517,
		"VNG":                   322,
		"Dream11":               387,
		"RocketWerkz":           341,
		"Mayday":                360,
		"5 Lives":               489,
		"Mod.io":                407,
		"Uppercut":              417,
		"Bokeh":                 396,
		"Ubitus":                476,
		"Tookyo":                326,
		"Liona":                 363,
		"JP Games":              495,
		"Next Ninja":            478,
		"Eggvision":             154,
		"Taskey":                165,
		"Unseen":                180,
		"Toylogic":              422,
		"F4samurai":             458,
		"SG RPG":                335,
		"Shiftup":               337,
		"Netmarble F&C":         353,
		"Line Games":            364,
		"Flint":                 381,
		"433":                   405,
		"Two Hands":             324,
		"Table One":             328,
		"NYOU":                  350,
		"Ngel":                  351,
		"Neostream":             355,
		"Ikina":                 371,
		"Haegin":                374,
		"Floppy":                380,
		"Iggymob":               192,
		"NXN":                   157,
		"PU Productions":        521,
		"Facepunch":             529,
		"Novarama":              469,
		"Gruby":                 460,
		"1C Game":               540,
		"Artisan":               435,
		"Sunny Side Up":         446,
		"Fromsoftware":          552,
		"Reikon":                509,
		"Red Wheel":             561,
		"Lighthouse":            693,
		"Raid Base":             699,
		"Sunny Side Up":         446,
		"Redlab":                156,
		"Mistil":                561,
		"Crytek":                142,
		"Bespoke Pixel":         753,
		"Digital Confectioners": 501,
		"库洛":                    667,
		"迷你玩":                   252,
		"银汉":                    295,
		"朱雀网络":                  483,
		"广州深蓝互动":                262,
		"因陀罗（帝释天":               270,
		"柳叶刀":                   487,
		"广州魂起":                  186,
		"明昼科技":                  204,
		"左东东（彩蛋）":               263,
		"珠海黑沙":                  408,
		"Doracat多罗猫":            161,
		"玄星":                    297,
		"纳仕":                    249,
		"谷得":                    264,
		"宇宙罐":                   289,
		"游戏科学":                  290,
		"无端科技":                  301,
		"乐府游戏":                  245,
		"灵刃":                    257,
		"欢乐互娱":                  484,
		"上海梦求":                  160,
		"宝可拉":                   309,
		"奇侠":                    176,
		"上海鹿游":                  254,
		"乐曼多":                   481,
		"成都星柒玩":                 155,
		"Tulip（那朵花）":            188,
		"龙渊游戏":                  480,
		"上海沐睦":                  492,
		"钛核网络":                  242,
		"帕斯亚":                   247,
		"Facet（上海琢面）":           526,
		"上海游牛":                  173,
		"润梦":                    311,
		"森霆":                    310,
		"四维八方":                  434,
		"91Act（成都格斗）":           278,
		"上海晨游":                  273,
		"阿哇龙":                   499,
		"苏州天魂":                  488,
		"胖布丁":                   316,
		"暗星":                    246,
		"零犀":                    256,
		"Indie（独立开发）":           277,
		"杭州灵羽":                  139,
		"云南盒子怪":                 410,
		"上海月之底":                 286,
		"成都擎龙":                  164,
		"世纪华通/盛趣	":              307,
		"英雄互娱":                  692,
		"攸乐":                    292,
		"西山居":                   300,
		"祖龙":                    279,
		"北京云畅":                  285,
		"灵游坊":                   505,
		"元趣":                    287,
		"水果堂":                   306,
		"像素":                    299,
		"所思":                    304,
		"星合互娱":                  241,
		"星海互娱":                  298,
		"卡布姆":                   539,
		"拱顶石":                   265,
		"玩心":                    303,
		"共感之脑":                  420,
		"龙拳风暴":                  255,
		"破晓互动":                  315,
		"华清飞扬":                  260,
		"酷酷跑":                   190,
		"中清龙图":                  281,
		"七号笔迹":                  314,
		"天津队友":                  268,
		"威魔纪元":                  181,
		"万象皆春":                  194,
		"散爆":                    479,
		"番糖":                    267,
		"黑桃":                    485,
		"游戏公国	":                 433,
		"掌梦":                    486,
		"莫彼吾斯":                  201,
		"苏州五十一区":                519,
		"阿佩吉":                   276,
		"氩紫":                    296,
		"紫月":                    280,
		"朝露":                    261,
		"鬼脸":                    174,
		"影之月":                   294,
		"猫布丁":                   253,
		"凡帕斯":                   266,
		"吉艾斯球":                  162,
		"波波":                    274,
		"彼方":                    195,
		"紫辰":                    702,
		"当乐":                    271,
		"任玩堂":                   312,
		"壹多互娱":                  542,
		"百奥":                    275,
		"青瓷":                    172,
		"量子动力":                  462,
		"十勇士":                   437,
		"杭州不鸣（Nanshan Booming）": 182,
		"乐元素":                   258,
		"叠纸":                    520,
		"隆匠":                    308,
	}
}
