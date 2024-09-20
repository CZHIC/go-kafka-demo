package msg

var DataBrainInfo string
var StudioList map[string]int
var ProductList map[int]int
var JpList []int // 日本公司

var StudioListTest map[string]int
var JpListTest []int // 日本公司
var ProductListTest map[int]int

func init() {

	StudioListTest = map[string]int{
		"3022b1c1-d6b2-475d-9765-95c8712a180c": 707,
		"3cae090b-ed2d-95f8-79a9-e32ca480258f": 920,
	}
	JpListTest = []int{920}

	ProductListTest = map[int]int{
		12: 49,
		13: 51,
		23: 52,
		22: 507,
		33: 920,
	}
	ProductList = map[int]int{
		15: 826,
		16: 827,
		17: 811,
		32: 393,
		6:  41,
		24: 58,
		30: 62,
		36: 226,
		58: 2982,
		11: 46,
		12: 47,
		21: 55,
		19: 53,
		2:  38,
		1:  36,
		9:  2551,
		77: 246,
		59: 2983,
		3:  39,
		4:  334,
		78: 3,
		18: 812,
		75: 2986,
		7:  2987,
		68: 2539,
		74: 2988,
		79: 2989,
		46: 2562,
		80: 2990,
		81: 2991,
		41: 67,
		45: 71,
		82: 2981,
		89: 3187,
		84: 3188,
		86: 3189,
		83: 3190,
		90: 3192,
		85: 3193,
		88: 3191,
		87: 3194,
	}

	StudioList = map[string]int{
		"3022b1c1-d6b2-475d-9765-95c8712a180c": 404,
		"70a4451e-3e12-ac31-2553-051ec6b0f9b2": 390,
		"80b39a2b-6b55-4066-868b-029bbd2bc8b1": 383,
		"24e4201e-8ac9-1488-62b8-c9ecf483c6b5": 378,
		"c522b63b-cd72-4cc1-8454-71337edd0db1": 375,
		"7e15eef8-c655-4928-a8ff-3a2662bf53c7": 368,
		"96c03483-b8c6-4fd9-a25b-2215e4b7a9f1": 338,
		"4306c6a0-b2f6-4194-3896-fabd3350d103": 333,
		"c69f9642-bda6-4347-92e7-e0d33d27fd00": 317,
		"d24c18d1-7600-395e-55a3-52648c41a1d5": 331,
		"1a195f56-ffe9-c556-14f1-4f731b44f300": 330,
		"bokehgamestudio":                      396,
		"f0b19907-a6fe-5d0d-c44a-9e3db9ddb256": 593,
		"579787ca-907c-f5ae-9567-7b99a6c39039": 472,
		"8c84ac3e-ace4-4ba0-86e5-dee26105b99a": 143,
		"unseen-inc-tokyo":                     180,
		"87629a56-0325-11ec-bc3d-525400d129b6": 477,
		"31231154-e6a8-0f09-d5db-615a874639f2": 370,
		"ca3001c7-2af3-4498-8296-a33ff83da153": 465,
		"ffbb640e-8df4-4d22-944a-1633b217c2b4": 346,
		"cc69ad75-2fa2-a3f5-5be2-40fa618f141b": 547,
		"49e40671-2c57-7c3f-46ce-f615ae7b53c6": 358,
		"f3663ae6-6c2c-4cfc-98b9-305a12106236": 325,
		"b9eb2822-306d-491d-99a8-45dbf24e4e27": 581,
		"e14d50fc-f96e-11ec-b848-52540061e6c5": 514,
		"fcf4f7b0-b5b6-4f63-b7c5-194ecb536e68": 656,
		"36e6b647-7933-0f41-30b7-859ae5861364": 698,
		"a5046e37-9d89-fd14-069a-1df7924d66ed": 142,
		"2c8511e0-d4c1-11ec-9bb0-525400d129b6": 183,
		"yx3001c7-2af3-4498-8296-a33ff83da153": 293,
	}

	JpList = []int{
		346,
		477,
		656,
		180,
	}

	DataBrainInfo = `
	{
    "code": 200,
    "msg": "ok",
    "data": [
        {
            "detail": {
                "project_name": "Kingdom Rush 5",
                "studio_id": "31231154-e6a8-0f09-d5db-615a874639f2",
                "platform": [
                    "mobile",
                    "pc"
                ],
                "studio": "Ironhide Game",
                "project_id": "63",
                "priority": "Low",
                "genre": [
                    "Casual"
                ],
                "expanded_genre": "Strategy",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Unknown",
                "ip": "Existing",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20240630",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pre-pitch",
                        "milestone_node": "",
                        "date": "20220901",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "63",
                        "child": null
                    },
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20221031",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "63",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20240630",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "63",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "julienbares",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 366,
                                    "sort": 0
                                },
                                {
                                    "name": "Jamey",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "",
                                    "id": 436,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Project MAD",
                "studio_id": "31231154-e6a8-0f09-d5db-615a874639f2",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Ironhide Game",
                "project_id": "65",
                "priority": "Low",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20241231",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20220901",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "65",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20241231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "65",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "julienbares",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 369,
                                    "sort": 0
                                },
                                {
                                    "name": "Howardtang",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4005",
                                    "id": 439,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "JUNKWORLD",
                "studio_id": "31231154-e6a8-0f09-d5db-615a874639f2",
                "platform": [
                    "pc",
                    "mobile"
                ],
                "studio": "Ironhide Game",
                "project_id": "43",
                "priority": "Low",
                "genre": [
                    "Strategy"
                ],
                "expanded_genre": "STRATEGY, TOWER DEFENSE",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20231031",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Soft launch",
                        "milestone_node": "",
                        "date": "20220531",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "43",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20231031",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "43",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "julienbares",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 371,
                                    "sort": 0
                                },
                                {
                                    "name": "howardtang",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4005",
                                    "id": 441,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Kingdom Rush Battle",
                "studio_id": "31231154-e6a8-0f09-d5db-615a874639f2",
                "platform": [
                    "mobile",
                    "pc"
                ],
                "studio": "Ironhide Game",
                "project_id": "64",
                "priority": "Low",
                "genre": [
                    "Casual"
                ],
                "expanded_genre": "Strategy",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Unknown",
                "ip": "Existing",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20240630",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pre-pitch",
                        "milestone_node": "",
                        "date": "20221031",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "64",
                        "child": null
                    },
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20221031",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "64",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20240630",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "64",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "jasmineyywu",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "10543",
                                    "id": 302,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "julienbares",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 385,
                                    "sort": 0
                                },
                                {
                                    "name": "Howardtang",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4005",
                                    "id": 455,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Vampire The Masquerade: Bloodhunt",
                "studio_id": "96c03483-b8c6-4fd9-a25b-2215e4b7a9f1",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Sharkmob",
                "project_id": "17",
                "priority": "Low",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "Battle royale",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Self publish",
                "ip": "Licensed",
                "tagline": "“Vampires suck blood from human characters and non-hostile subjects. Your character can use vampiric powers, weapons and wit to eradicate your enemies and deal with the hunters.”",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "e63c73107ebcc41248eb4c54baa975e38",
                        "unified_id": "e63c73107ebcc41248eb4c54baa975e38",
                        "name": "Midnight Ghost Hunt",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/3f5d20554c57156c1dd78be2ebbead3e.jpg",
                        "game_type": 2,
                        "genre": "Action Games|Shooter|First-Person Shooter|Action|Indie",
                        "sub_genre": "",
                        "country_en_name": "Sweden",
                        "country_en_name_abbr": "SE",
                        "region": "Europe",
                        "publisher_name": "Coffee Stain Publishing",
                        "developer_name": "Vaulted Sky Games"
                    },
                    {
                        "id": "ufc454d9b1af70b40588e2a6fa4da4a8b",
                        "unified_id": "ufc454d9b1af70b40588e2a6fa4da4a8b",
                        "name": "PUBG MOBILE",
                        "cover": "https://ogdb-cdn.intlgame.cn/pic_by_handle/pubg.png",
                        "game_type": 1,
                        "genre": "Shooting",
                        "sub_genre": "Battle Royale",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "Rekoo Japan|HotCool Game|Row Business Solutions|VNG|Krafton|Tencent",
                        "developer_name": "Rekoo Japan|HotCool Game|Row Business Solutions|VNG|Krafton|Tencent"
                    }
                ],
                "current_headcounts": "",
                "launch_date": "20220427",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20210701",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "17",
                        "child": null
                    },
                    {
                        "milestone": "Soft launch",
                        "milestone_node": "",
                        "date": "20210907",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "17",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20211201",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "17",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20220427",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "17",
                        "child": [
                            {
                                "milestone": "Decision gate",
                                "milestone_node": "",
                                "date": "20220831",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "17"
                            }
                        ]
                    },
                    {
                        "milestone": "Game Sunset",
                        "milestone_node": "",
                        "date": "20221031",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "17",
                        "child": null
                    },
                    {
                        "milestone": "Game shutdown ",
                        "milestone_node": "",
                        "date": "20230731",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "17",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 395,
                                    "sort": 0
                                },
                                {
                                    "name": "tomoconnor",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1035",
                                    "id": 466,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 500,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "eric",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "",
                                    "id": 313,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "IRON MARINES 2",
                "studio_id": "31231154-e6a8-0f09-d5db-615a874639f2",
                "platform": [
                    "pc",
                    "mobile"
                ],
                "studio": "Ironhide Game",
                "project_id": "42",
                "priority": "Low",
                "genre": [
                    "Strategy"
                ],
                "expanded_genre": "REAL TIME STRATEGY",
                "business_model": [
                    "P+MTX"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Unknown",
                "ip": "Sequel",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20220930",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "First playable",
                        "milestone_node": "",
                        "date": "20220513",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "42",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20220601",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "42",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20220930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "42",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "julienbares",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 370,
                                    "sort": 0
                                },
                                {
                                    "name": "howardtang",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4005",
                                    "id": 440,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "V Rising",
                "studio_id": "d24c18d1-7600-395e-55a3-52648c41a1d5",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Stunlock Studios",
                "project_id": "24",
                "priority": "Low",
                "genre": [
                    "RPG",
                    "SOC",
                    "Action"
                ],
                "expanded_genre": "SOC RPG",
                "business_model": [
                    "P+DLC"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Stunlock Studios (+Yooreka in Asia)",
                "ip": "New",
                "tagline": "Awaken as a vampire. Hunt for blood in nearby settlements to regain your strength and evade the scorching sun to survive. Raise your castle and thrive in an ever-changing open world full of mystery. Gain allies online and conquer the land of the living.",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "e5bd811c75706e21353bc3bceb92215c7",
                        "unified_id": "e5bd811c75706e21353bc3bceb92215c7",
                        "name": "Valheim",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/c0d1cca7b5574981d6e7acdaebf750ba.jpg",
                        "game_type": 2,
                        "genre": "Adventure|Action Games|Action-Adventure|RPG|Role-playing (RPG)|Action|Action & adventure|Survival|动作与冒险|Azione e avventura|Action et aventure|アクション & アドベンチャー|Indie|Acción y aventura|Actie en avontuur",
                        "sub_genre": "",
                        "country_en_name": "Sweden",
                        "country_en_name_abbr": "SE",
                        "region": "Europe",
                        "publisher_name": "Coffee Stain Publishing",
                        "developer_name": "Iron Gate AB"
                    },
                    {
                        "id": "e53e156fd95b67e61a1c302842aa9d145",
                        "unified_id": "e53e156fd95b67e61a1c302842aa9d145",
                        "name": "Battlerite",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/da337d2d486becdca7bdff707a7017d0.jpg",
                        "game_type": 2,
                        "genre": "Battle Arena|Action|MOBA|Action Games",
                        "sub_genre": "",
                        "country_en_name": "Sweden",
                        "country_en_name_abbr": "SE",
                        "region": "Europe",
                        "publisher_name": "Stunlock Studios",
                        "developer_name": "Stunlock Studios"
                    }
                ],
                "current_headcounts": "35",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20211201",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "24",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20220517",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "24",
                        "child": null
                    },
                    {
                        "milestone": "Major Update PC",
                        "milestone_node": "",
                        "date": "20230517",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "24",
                        "child": null
                    },
                    {
                        "milestone": "Console Launch",
                        "milestone_node": "",
                        "date": "20240401",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "24",
                        "child": null
                    },
                    {
                        "milestone": "PC Full Launch",
                        "milestone_node": "",
                        "date": "20240517",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "24",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "sdecroix",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 376,
                                    "sort": 0
                                },
                                {
                                    "name": "mberthou",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "17097",
                                    "id": 446,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "sstellagan",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "4024",
                                    "id": 559,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Rickard Frisegård",
                                    "position": "CEO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 508,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "publishing",
                            "list": [
                                {
                                    "name": "Johan Ilves",
                                    "position": "Marketing Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 509,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Warframe Mobile",
                "studio_id": "70a4451e-3e12-ac31-2553-051ec6b0f9b2",
                "platform": [
                    "mobile"
                ],
                "studio": "Digital Extremes",
                "project_id": "2",
                "priority": "Key",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "looter/shooter",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Digital Extremes",
                "ip": "Existing",
                "tagline": "",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "ed09fe5188fb2059d33eb6ba04958bbb8",
                        "unified_id": "ed09fe5188fb2059d33eb6ba04958bbb8",
                        "name": "Tom Clancy's The Division",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/adf676d00f4dc27d82de8da0da881c63.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Action Games|Tactical Third Person Shooter|Action|Third-Person Shooter|Survival",
                        "sub_genre": "",
                        "country_en_name": "France",
                        "country_en_name_abbr": "FR",
                        "region": "Europe",
                        "publisher_name": "Ubisoft",
                        "developer_name": "Massive Entertainment"
                    }
                ],
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Closed Beta",
                        "milestone_node": "",
                        "date": "20221206",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "2",
                        "child": null
                    },
                    {
                        "milestone": "Full Launch",
                        "milestone_node": "",
                        "date": "20240612",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "2",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "celinaluo",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "583",
                                    "id": 305,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "francoischa",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "14648",
                                    "id": 388,
                                    "sort": 0
                                },
                                {
                                    "name": "faudet",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4020",
                                    "id": 458,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Sheldon Carter",
                                    "position": "pteam_contact",
                                    "link": "https://www.linkedin.com/in/sheldoncarter/",
                                    "user_id": "",
                                    "id": 522,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Den of Wolves",
                "studio_id": "3022b1c1-d6b2-475d-9765-95c8712a180c",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "10 Chambers Collective",
                "project_id": "1",
                "priority": "Top 10",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "Co-op PVE shooter",
                "business_model": [
                    "Premium",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "10 Chambers Collective",
                "ip": "New",
                "tagline": "“Take on the role of a criminal in an unregulated corporate haven in Midway City”\nA session-based PvE Coop heist shooter with a strong narrative. \"Payday meets Destiny\". Build as GAAS from the ground up",
                "ambition": "A co-op PVE heist fantasy",
                "similar_games": [
                    {
                        "id": "e11000000288",
                        "unified_id": "e11000000288",
                        "name": "Destiny 2",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/e7a0f4301d1f88cbf9f79296cbca7c9a.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Action Games|First-Person Shooter|RPG|Role-playing (RPG)|Action|MMO|Action & adventure|FPS|Azione e avventura|Action et aventure|アクション & アドベンチャー|Acción y aventura",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Bungie",
                        "developer_name": "Bungie"
                    },
                    {
                        "id": "e1ef60cef2767697a67e179a32c0c00f4",
                        "unified_id": "e1ef60cef2767697a67e179a32c0c00f4",
                        "name": "Payday 2",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/7d516edb8d6ad5626b59e670f9933c76.jpg",
                        "game_type": 3,
                        "genre": "First-Person Shooter|Tactical|FPS|Shooter|Role-playing (RPG)",
                        "sub_genre": "",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "Starbreeze Studios|505 Games",
                        "developer_name": "Overkill Software"
                    },
                    {
                        "id": "ea597a0cc3bb3b71c9c281129b9a51f49",
                        "unified_id": "ea597a0cc3bb3b71c9c281129b9a51f49",
                        "name": "Grand Theft Auto",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/dfbac39fe6cfb8415677215091a063c0.jpg",
                        "game_type": 2,
                        "genre": "Action Games|Action|Open-World Action",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Rockstar Games",
                        "developer_name": "Rockstar North"
                    },
                    {
                        "id": "eda9a90de4243ad327a5bf835998e6ccb",
                        "unified_id": "eda9a90de4243ad327a5bf835998e6ccb",
                        "name": "PAYDAY 3",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/f4fb0c1516191f561b55041227995cc3.jpg",
                        "game_type": 2,
                        "genre": "Action Games|Shooter|FPS|Sparatutto|Action|First-Person Shooter|Schieten|Jeux de tir|シューティング",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Starbreeze Studios"
                    }
                ],
                "current_headcounts": "87",
                "launch_date": "20251231",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20220427",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "1",
                        "child": null
                    },
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20221214",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "1",
                        "child": null
                    },
                    {
                        "milestone": "Internal PB7 Milestone",
                        "milestone_node": "",
                        "date": "20231130",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "1",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20240531",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "1",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20241118",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "1",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20251231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "1",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "sdecroix",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 372,
                                    "sort": 0
                                },
                                {
                                    "name": "olegg",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1022",
                                    "id": 442,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "gracehhuang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "11050",
                                    "id": 519,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "publishing",
                            "list": [
                                {
                                    "name": "Ulf Andersson",
                                    "position": "CEO / Creative Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 110,
                                    "sort": 0
                                },
                                {
                                    "name": "Svante Vinternatt",
                                    "position": "COO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 111,
                                    "sort": 0
                                },
                                {
                                    "name": "Chris Tattersall",
                                    "position": "CMO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 112,
                                    "sort": 0
                                },
                                {
                                    "name": "Oscar J-T Holm",
                                    "position": "CSO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 113,
                                    "sort": 0
                                },
                                {
                                    "name": "Sam Millen",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 114,
                                    "sort": 0
                                },
                                {
                                    "name": "Hjalmar Vikström",
                                    "position": "Game Design Director and CTO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 116,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Project Derby",
                "studio_id": "1a195f56-ffe9-c556-14f1-4f731b44f300",
                "platform": [
                    "pc"
                ],
                "studio": "Sumo Digital",
                "project_id": "59",
                "priority": "Middle",
                "genre": [
                    "MMO"
                ],
                "expanded_genre": "Racing",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Sumo Digital",
                "publisher": "Level Infinite",
                "ip": "New",
                "tagline": "Kart racing; Battle Royale",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "e5399566b81c5972aaf8f964e65bb13b2",
                        "unified_id": "e5399566b81c5972aaf8f964e65bb13b2",
                        "name": "Fall Guys",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/2fbc7d92d09e0270ea6fc2e8304f8627.jpg",
                        "game_type": 2,
                        "genre": "Racing|Sport|Platformer|Platform|Action|Massively Multiplayer Games|MMO|Party|Indie|Arcade",
                        "sub_genre": "",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "Devolver Digital|Epic Games",
                        "developer_name": "Mediatonic"
                    },
                    {
                        "id": "u54fd3364f7949f528a6751f477279df0",
                        "unified_id": "u54fd3364f7949f528a6751f477279df0",
                        "name": "KartRider Rush+",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/fbffb61e40cfc96491948795f4dc0dda.jpg",
                        "game_type": 1,
                        "genre": "Racing",
                        "sub_genre": "Competitive Racing",
                        "country_en_name": "Japan",
                        "country_en_name_abbr": "JP",
                        "region": "Japan",
                        "publisher_name": "NEXON",
                        "developer_name": "NEXON"
                    }
                ],
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20220922",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "59",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20221213",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "59",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20231231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "59",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20240131",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "59",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20240331",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "59",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 334,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 416,
                                    "sort": 0
                                },
                                {
                                    "name": "tomoconnor",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1035",
                                    "id": 495,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 496,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Wayfinder",
                "studio_id": "70a4451e-3e12-ac31-2553-051ec6b0f9b2",
                "platform": [
                    "mobile",
                    "pc",
                    "console"
                ],
                "studio": "Digital Extremes",
                "project_id": "4",
                "priority": "Low",
                "genre": [
                    "RPG/Action"
                ],
                "expanded_genre": "ARPG",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "WFH",
                "publisher": "TBC",
                "ip": "New",
                "tagline": "",
                "ambition": "Wayfinder is a blended social, co-op dungeon runner with Action RPG MMO \"lite\" open wolrld elements.  The game allows the player to choose and outfit a hero, upgrade and refine their playstyle, weapons, looks and level before moving on to alternate heroes to advance in a similar manner.  The world of Everon is an evergrowing experience that introduces new open world areas and new dungeons to explore as the game progresses over the years.\nWill launch on PC (Steam, Launcher, EGS), PS4/PS5, Xbox One/Xbox x/s, iOs, Android",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20240226",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20220228",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "4",
                        "child": null
                    },
                    {
                        "milestone": "Friends and Family test",
                        "milestone_node": "",
                        "date": "20221031",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "4",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta 1",
                        "milestone_node": "",
                        "date": "20221214",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "4",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta 2",
                        "milestone_node": "",
                        "date": "20230131",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "4",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta 3",
                        "milestone_node": "",
                        "date": "20230331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "4",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20230828",
                        "date_confirm": "2",
                        "gate_review": "Redo",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "4",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20240226",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "4",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "francoischa",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "14648",
                                    "id": 387,
                                    "sort": 0
                                },
                                {
                                    "name": "faudet",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4020",
                                    "id": 457,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "Murphy Pettypiece",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "",
                                    "id": 304,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Warframe",
                "studio_id": "70a4451e-3e12-ac31-2553-051ec6b0f9b2",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Digital Extremes",
                "project_id": "78",
                "priority": "Top 5",
                "genre": [
                    "MMO"
                ],
                "expanded_genre": "PVE",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Digital Extremes",
                "publisher": "Digital Extremes",
                "ip": "Existing",
                "tagline": "Ninja Play Free",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "291",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Live : Whisper in the Wall ",
                        "milestone_node": "",
                        "date": "20231221",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "78",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "francoischa",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "14648",
                                    "id": 526,
                                    "sort": 0
                                },
                                {
                                    "name": "faudet",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4020",
                                    "id": 527,
                                    "sort": 0
                                },
                                {
                                    "name": "celinaluo",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "583",
                                    "id": 528,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Rebecca Ford",
                                    "position": "Creative Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 529,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "D2 - Bachman",
                "studio_id": "1a195f56-ffe9-c556-14f1-4f731b44f300",
                "platform": [
                    "mobile",
                    "pc",
                    "console"
                ],
                "studio": "Sumo Digital",
                "project_id": "71",
                "priority": "Low",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Unknown",
                "ip": "Existing",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20240930",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20221122",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "71",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20221212",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "71",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20231120",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "71",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20240630",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "71",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "71",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "71",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 417,
                                    "sort": 0
                                },
                                {
                                    "name": "tomoconnor",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1035",
                                    "id": 498,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 499,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 335,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Dying Light 2",
                "studio_id": "36e6b647-7933-0f41-30b7-859ae5861364",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Techland",
                "project_id": "80",
                "priority": "Key",
                "genre": [
                    "RPG",
                    "Action"
                ],
                "expanded_genre": "OpenWorld, Survival",
                "business_model": [
                    "Premium",
                    "Downloadle Content",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "Techland",
                "publisher": "Techland",
                "ip": "Existing",
                "tagline": "- Find yourself in the middle of unraveling zombie apocalypse,\u000bamidst chaos and destruction.\n- Get ready, either solo or with a group of friends, to stand against the relentless enemy that knows no mercy.\n- Enjoy the freedom of exploration, movement & combat to play the way you want. \n- Mastered your zombie slaughtering skills during the day. Now try surviving the night.",
                "ambition": "DL franchise: Create Ultimate Zombie Experience",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20220204",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20220204",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "80",
                        "child": null
                    },
                    {
                        "milestone": "2nd anniversary",
                        "milestone_node": "",
                        "date": "20240204",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "80",
                        "child": null
                    },
                    {
                        "milestone": "DLC2",
                        "milestone_node": "",
                        "date": "20240613",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "80",
                        "child": null
                    },
                    {
                        "milestone": "Dying Light: Beyond Human DLC",
                        "milestone_node": "",
                        "date": "20241113",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "80",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Paweł Marchewka",
                                    "position": "CEO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 549,
                                    "sort": 0
                                },
                                {
                                    "name": "Adrian Ciszewski",
                                    "position": "Chief Creative Officer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 550,
                                    "sort": 0
                                },
                                {
                                    "name": "Oleg Klapovskiy",
                                    "position": "Chief Publishing Officer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 551,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "julienbares",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "707",
                                    "id": 547,
                                    "sort": 0
                                },
                                {
                                    "name": "tonyxchen",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "12005",
                                    "id": 548,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Dying Light: World",
                "studio_id": "36e6b647-7933-0f41-30b7-859ae5861364",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Techland",
                "project_id": "81",
                "priority": "",
                "genre": [
                    "RPG",
                    "Action"
                ],
                "expanded_genre": "OpenWorld, Online",
                "business_model": [
                    "Microtransaction",
                    "Premium",
                    "Downloadle Content"
                ],
                "image_name": "",
                "developer": "Techland",
                "publisher": "Techland",
                "ip": "Existing",
                "tagline": "- Bridging Solo & Multiplayer: PVE first, Solo compatible\n- Freedom of exploration & movement\n- Visceral combat with emphasis on elemental powers\n- Strong Day & Night cycle + risk/reward extraction loop increasing the tension\n- Transition to multiplayer (built for coop + asynchronous social interactions)\n- New game structure tailored for GAAS  built for long lasting engagement",
                "ambition": "Building upon and improving what makes Dying Light a great and unique experience, mixed with proven GAAS ingredients.",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "MVP 0.5",
                        "milestone_node": "",
                        "date": "20241218",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "81",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "julienbares",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "707",
                                    "id": 552,
                                    "sort": 0
                                },
                                {
                                    "name": "tonyxchen",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "12005",
                                    "id": 553,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Paweł Marchewka",
                                    "position": "CEO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 554,
                                    "sort": 0
                                },
                                {
                                    "name": "Oleg Klapovskiy",
                                    "position": "Chief Publishing Officer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 555,
                                    "sort": 0
                                },
                                {
                                    "name": "Adrian Ciszewski",
                                    "position": "Chief Creative Officer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 556,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Path of Exile 2",
                "studio_id": "c522b63b-cd72-4cc1-8454-71337edd0db1",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Grinding Gear Games",
                "project_id": "11",
                "priority": "Top 5",
                "genre": [
                    "RPG/Action"
                ],
                "expanded_genre": "ARPG",
                "business_model": [
                    "F2P, GaaS"
                ],
                "image_name": "",
                "developer": "Project finance",
                "publisher": "Self publish",
                "ip": "Existing",
                "tagline": "Path of Exile 2 is a new six-act storyline that is available alongside the original Path of Exile 1 campaign. Path of Exile 2 leverages much of the existing content for Path of Exile so that at launch, it has significant additional content ready to go.",
                "ambition": "AAA sequel of Path of Exile",
                "similar_games": [
                    {
                        "id": "e952663fcf495e58954fe9922f4a837a0",
                        "unified_id": "e952663fcf495e58954fe9922f4a837a0",
                        "name": "Diablo III",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/dfe8f6e47c6914c32aced7a6f0f972fb.jpg",
                        "game_type": 3,
                        "genre": "Adventure|Role Playing|Hack and slash/Beat 'em up|Role-Playing|RPG|Role-playing (RPG)|Action|Action RPG|Party",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Blizzard Entertainment",
                        "developer_name": "Blizzard Entertainment"
                    },
                    {
                        "id": "ecbbb35ac3bd34bec0a904ceb135d8ed0",
                        "unified_id": "ecbbb35ac3bd34bec0a904ceb135d8ed0",
                        "name": "Diablo IV",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/9c26eeb5b20fb4dbbabe074d2651707e.jpg",
                        "game_type": 3,
                        "genre": "Adventure|Role Playing|Hack and slash/Beat 'em up|RPG|Role-playing (RPG)|Action|Action RPG",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Blizzard Entertainment",
                        "developer_name": "Blizzard Entertainment"
                    }
                ],
                "current_headcounts": "102",
                "launch_date": "20241230",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20210929",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "11",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20220921",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "11",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20240329",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "11",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta",
                        "milestone_node": "",
                        "date": "20240607",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "11",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20241230",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "11",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "bubuqian",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 418,
                                    "sort": 0
                                },
                                {
                                    "name": "reginagao",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "316",
                                    "id": 489,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "Celinaluo",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "583",
                                    "id": 336,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Chris Wilson",
                                    "position": "CEO & managing director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 107,
                                    "sort": 0
                                },
                                {
                                    "name": "Jonathan Rogers",
                                    "position": "Game director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 108,
                                    "sort": 0
                                },
                                {
                                    "name": "Trevor Gamon",
                                    "position": "Producer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 109,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "B4B2",
                "studio_id": "f3663ae6-6c2c-4cfc-98b9-305a12106236",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Turtle Rock Studios",
                "project_id": "68",
                "priority": "Top 5",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "",
                "business_model": [
                    "Premium",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "TRS",
                "publisher": "Turtle Rock Studios",
                "ip": "Existing",
                "tagline": "",
                "ambition": "AAA, GaaS, #1 in co-op zombie PVE shooter genre.",
                "similar_games": [
                    {
                        "id": "e161ae8c05503db74c889dc48d3882d89",
                        "unified_id": "e161ae8c05503db74c889dc48d3882d89",
                        "name": "Dead Island 2",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/676979f17e1fb9f114a7581fd1119991.jpg",
                        "game_type": 3,
                        "genre": "Action|Horror|FPS|Shooter|Role-playing (RPG)|Adventure|Hack and slash/Beat 'em up",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Deep Silver Dambuster Studios"
                    },
                    {
                        "id": "e2eb0747e41fed2f481656a6d5d636b54",
                        "unified_id": "e2eb0747e41fed2f481656a6d5d636b54",
                        "name": "Left 4 Dead",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/7a66b82d29431bb8511ada35d970eef2.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Action Games|FPS",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Valve",
                        "developer_name": "Valve"
                    },
                    {
                        "id": "e55a84bf7f676f88637f2a04c0db42675",
                        "unified_id": "e55a84bf7f676f88637f2a04c0db42675",
                        "name": "Left 4 Dead",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/1f6d2150557a886fede665b46f4d02f0.jpg",
                        "game_type": 3,
                        "genre": "FPS|Shooter",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Valve",
                        "developer_name": "Valve|Certain Affinity|Turtle Rock Studios"
                    },
                    {
                        "id": "e6b4f5ee702dd3943ecb3ba2812312e69",
                        "unified_id": "e6b4f5ee702dd3943ecb3ba2812312e69",
                        "name": "Escape Dead Island",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/3a1a74ca8cb3874098a12f0434cb0197.jpg",
                        "game_type": 3,
                        "genre": "Survival|Adventure|Action|Simulator|Indie",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Fatshark"
                    },
                    {
                        "id": "e72b4feda779c36a6b6d3cb5eb8efe5d1",
                        "unified_id": "e72b4feda779c36a6b6d3cb5eb8efe5d1",
                        "name": "Warhammer 40,000: Darktide",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/3553b1ed867da608a04ed7599dcf7377.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Action Games|Role-playing (RPG)|Action|Action & adventure|FPS|Azione e avventura|Action et aventure|アクション & アドベンチャー|Strategy|Indie|Acción y aventura|Actie en avontuur",
                        "sub_genre": "",
                        "country_en_name": "Sweden",
                        "country_en_name_abbr": "SE",
                        "region": "Europe",
                        "publisher_name": "Fatshark",
                        "developer_name": "Fatshark"
                    },
                    {
                        "id": "ebdda384fc5ff86d75fbf8d42ebe5798f",
                        "unified_id": "ebdda384fc5ff86d75fbf8d42ebe5798f",
                        "name": "Dead Island: Definitive Edition",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/dad420c776ebc74e7abf0e2e4f586491.jpg",
                        "game_type": 3,
                        "genre": "Action|Survival|Horror|Shooter|Compilation|Role-playing (RPG)|Adventure|Hack and slash/Beat 'em up",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Techland"
                    },
                    {
                        "id": "e144f730dc2d687b3c8e852da92a7432d",
                        "unified_id": "e144f730dc2d687b3c8e852da92a7432d",
                        "name": "Left 4 Dead 2",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/2f89bab26b471a74c7730b4e5e24095b.jpg",
                        "game_type": 3,
                        "genre": "Shooter|Tactical FPS",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Valve",
                        "developer_name": "Turtle Rock Studios|Valve"
                    },
                    {
                        "id": "e203ab9444a5f0347892fb4a5f011e48b",
                        "unified_id": "e203ab9444a5f0347892fb4a5f011e48b",
                        "name": "Dead Island: Riptide - Definitive Edition",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/1c61fa374e768fb5238f6d026f0ea7d8.jpg",
                        "game_type": 3,
                        "genre": "Action|Misc|FPS|Shooter|Role-playing (RPG)",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Techland"
                    },
                    {
                        "id": "e94ed78ece41e234cb3afd4b5c88ed865",
                        "unified_id": "e94ed78ece41e234cb3afd4b5c88ed865",
                        "name": "Dead Island: Epidemic",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/894e455512b0701bb4c4fa6eece620fa.jpg",
                        "game_type": 2,
                        "genre": "MOBA",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Stunlock Studios"
                    },
                    {
                        "id": "e756bff6af3138c5a58da0965a177c9c6",
                        "unified_id": "e756bff6af3138c5a58da0965a177c9c6",
                        "name": "Dead Island 2",
                        "cover": "https://ogdb-cdn.intlgame.cn/pic_by_handle/pc/dead_island2.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Horror|Hack and slash/Beat 'em up|Role-playing (RPG)|Action|FPS",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Deep Silver Dambuster Studios"
                    },
                    {
                        "id": "e5d639b0a893344937cbddfaaca7094e0",
                        "unified_id": "e5d639b0a893344937cbddfaaca7094e0",
                        "name": "Dead Island Retro Revenge",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/06462dbdeed105b8df1762be05ab59d3.jpg",
                        "game_type": 3,
                        "genre": "2D Beat-'Em-Up|Horror|Hack and slash/Beat 'em up|Fighting|Arcade",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Empty Clip Studios"
                    },
                    {
                        "id": "e3f7b9b7608ca0aac42c5ceeed5b71d41",
                        "unified_id": "e3f7b9b7608ca0aac42c5ceeed5b71d41",
                        "name": "Dead Island Definitive Edition",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/8553bd4eab5f9f9ec61cc19527c67033.jpg",
                        "game_type": 2,
                        "genre": "Adventure|Simulator|Survival|Action Games|Adventure Games|Shooter|Horror|Action|Hack and slash/Beat 'em up|Indie",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Techland"
                    },
                    {
                        "id": "e4637572ee6f3eb6a1b9fa1523451e680",
                        "unified_id": "e4637572ee6f3eb6a1b9fa1523451e680",
                        "name": "Dead Island: Bloodbath Arena",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/6ac64cf5b0f0cdd9b8915a7b64214f41.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Action|Survival",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Techland"
                    },
                    {
                        "id": "ea5f491a905834f466ba7412b63ee738c",
                        "unified_id": "ea5f491a905834f466ba7412b63ee738c",
                        "name": "DarkTide",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/c2699ae77a82a628734b1b3730eaf323.jpg",
                        "game_type": 2,
                        "genre": "Action Games",
                        "sub_genre": "",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "BlackDrop",
                        "developer_name": "BlackDrop"
                    },
                    {
                        "id": "u01265c57b87dfc091970e043f6fddd18",
                        "unified_id": "u01265c57b87dfc091970e043f6fddd18",
                        "name": "Redlight Survivor: Dead Island",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/eaab15c4cdcc6f51133f3543521a727c.jpg",
                        "game_type": 1,
                        "genre": "Action",
                        "sub_genre": "Others",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "OneSoft",
                        "developer_name": "OneSoft"
                    },
                    {
                        "id": "e5390379a2a0f497d056012855a5fd331",
                        "unified_id": "e5390379a2a0f497d056012855a5fd331",
                        "name": "Dead Island Retro Revenge",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/17f6a23ccdf8f1c5b4c1badee1d9217c.jpg",
                        "game_type": 2,
                        "genre": "2D Beat-'Em-Up|Action Games|Horror|Hack and slash/Beat 'em up|Arcade|Fighting",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Empty Clip Studios"
                    },
                    {
                        "id": "ea9db5d94930f0add249c2a5b6d1c998d",
                        "unified_id": "ea9db5d94930f0add249c2a5b6d1c998d",
                        "name": "Warhammer 40,000: Darktide",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/adbb88215ec8968e1e856ed9466be3b7.jpg",
                        "game_type": 3,
                        "genre": "Shooter|Adventure|Role-playing (RPG)|Action|FPS|Indie|Strategy",
                        "sub_genre": "",
                        "country_en_name": "Sweden",
                        "country_en_name_abbr": "SE",
                        "region": "Europe",
                        "publisher_name": "Fatshark",
                        "developer_name": "Fatshark"
                    },
                    {
                        "id": "ef7bd6cc8447632f76f5de1ae098cc98a",
                        "unified_id": "ef7bd6cc8447632f76f5de1ae098cc98a",
                        "name": "Dead Island: Riptide Definitive Edition",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/5499e7e18da6b8283519d036c6746d9f.jpg",
                        "game_type": 2,
                        "genre": "Role-playing (RPG)|Action Games|FPS|Shooter|Action",
                        "sub_genre": "",
                        "country_en_name": "Austria",
                        "country_en_name_abbr": "AT",
                        "region": "Europe",
                        "publisher_name": "Deep Silver",
                        "developer_name": "Techland"
                    }
                ],
                "current_headcounts": "164",
                "launch_date": "20260331",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20220930",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "68",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20230208",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "68",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20240630",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "68",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20250331",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "68",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20250630",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "68",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20251231",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "68",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20260331",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "68",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "gramxu",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 360,
                                    "sort": 0
                                },
                                {
                                    "name": "stanleyli",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "11515",
                                    "id": 504,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Soulframe",
                "studio_id": "70a4451e-3e12-ac31-2553-051ec6b0f9b2",
                "platform": [
                    "mobile",
                    "pc",
                    "console"
                ],
                "studio": "Digital Extremes",
                "project_id": "3",
                "priority": "Key",
                "genre": [
                    "RPG",
                    "Action"
                ],
                "expanded_genre": "Open world action/shooter",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Digital Extremes",
                "ip": "New",
                "tagline": "Evolved Warframe with more open map + MMO, richer environments, more melee combat and more mainstream fantasy theme",
                "ambition": "Soulframe is a free-to-play open-world adventure that is \"heavily influenced by themes of nature, restoration, and exploration\" alongside such works as Princess Mononoke and The NeverEnding Story. The conceit [in ‘Soulframe’] is that the world itself is a little angry about what’s been done to it, and the grounds underneath tend to shift throughout the day, So there’s going to be proceduralism within the cave networks and crevasses and so on underneath the world.",
                "similar_games": null,
                "current_headcounts": "20",
                "launch_date": "20250331",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20220513",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "3",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20230111",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "3",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20231205",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "3",
                        "child": null
                    },
                    {
                        "milestone": "Closed Alpha",
                        "milestone_node": "",
                        "date": "20231215",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "3",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20250331",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "3",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Sheldon Carter",
                                    "position": "Executive Producer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 143,
                                    "sort": 0
                                },
                                {
                                    "name": " Scott MacGregor",
                                    "position": "Design Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 145,
                                    "sort": 0
                                },
                                {
                                    "name": "Geoff Crookes",
                                    "position": "Art/Creative Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 146,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "francoischa",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "14648",
                                    "id": 389,
                                    "sort": 0
                                },
                                {
                                    "name": "faudet",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4020",
                                    "id": 459,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "celinaluo",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "583",
                                    "id": 306,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Dread Pilot",
                "studio_id": "7e15eef8-c655-4928-a8ff-3a2662bf53c7",
                "platform": [
                    "pc"
                ],
                "studio": "Klei Entertainment",
                "project_id": "16",
                "priority": "Low",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "2D space survival shooter",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Klei",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20230731",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "16",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "eddiechan",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 380,
                                    "sort": 0
                                },
                                {
                                    "name": "jsilveira",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "88",
                                    "id": 450,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "celinaluo",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "583",
                                    "id": 297,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Rotwood",
                "studio_id": "7e15eef8-c655-4928-a8ff-3a2662bf53c7",
                "platform": [
                    "pc"
                ],
                "studio": "Klei Entertainment",
                "project_id": "15",
                "priority": "Low",
                "genre": [
                    "RPG/Action"
                ],
                "expanded_genre": "Co-op 2D boss battler",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Klei",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20230331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "15",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "eddiechan",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 379,
                                    "sort": 0
                                },
                                {
                                    "name": "jsilveira",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "88",
                                    "id": 449,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "celinaluo",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "583",
                                    "id": 296,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Don't Starve Newhome",
                "studio_id": "7e15eef8-c655-4928-a8ff-3a2662bf53c7",
                "platform": [
                    "mobile"
                ],
                "studio": "Klei Entertainment",
                "project_id": "13",
                "priority": "Low",
                "genre": [
                    "SOC"
                ],
                "expanded_genre": "SOC",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Tencent",
                "ip": "Existing",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Mainland China Launch",
                        "milestone_node": "",
                        "date": "20221231",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "13",
                        "child": null
                    },
                    {
                        "milestone": "Global Launch",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "13",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "eddiechan",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 358,
                                    "sort": 0
                                },
                                {
                                    "name": "taihexu",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "906",
                                    "id": 425,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Seabass (Dune:Awakening)",
                "studio_id": "24e4201e-8ac9-1488-62b8-c9ecf483c6b5",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Funcom",
                "project_id": "9",
                "priority": "Top 10",
                "genre": [
                    "SOC",
                    "MMO"
                ],
                "expanded_genre": "MMOSOC, RPG, TPS",
                "business_model": [
                    "Premium",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Funcom",
                "ip": "Licensed",
                "tagline": "Abandoned on the most dangerous planet in the universe, you have to SURVIVE the desert, storms and sandworms; build and PROTECT your base; EXPAND through combat, crafting and spice harvesting, to become the one in CONTROL, as he who controls the spice, controls the universe",
                "ambition": "SOC MMO based on Dune IP",
                "similar_games": [
                    {
                        "id": "e6dbcff6a91065a23e1c406d9f1bddef3",
                        "unified_id": "e6dbcff6a91065a23e1c406d9f1bddef3",
                        "name": "Rust",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/5320f8e3380bd0e9211c4163f8699483.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Action Games|Action-Adventure|Role-playing (RPG)|Action|Survival|Indie",
                        "sub_genre": "",
                        "country_en_name": "United Kingdom",
                        "country_en_name_abbr": "GB",
                        "region": "Europe",
                        "publisher_name": "Facepunch Studios",
                        "developer_name": "Facepunch Studios"
                    },
                    {
                        "id": "efc7a5eee1d6ca668931bcebf69198b60",
                        "unified_id": "efc7a5eee1d6ca668931bcebf69198b60",
                        "name": "Conan Exiles",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/704a2d67c7dcadec76780dc1b958aec6.jpg",
                        "game_type": 2,
                        "genre": "Simulator|Adventure|Action Games|Action-Adventure|RPG|Simulation|Role-playing (RPG)|Action|RPG Games|MMO|Action & adventure|Survival|Fighting|Azione e avventura|Action et aventure|アクション & アドベンチャー|Strategy|Acción y aventura|Actie en avontuur",
                        "sub_genre": "",
                        "country_en_name": "Norway",
                        "country_en_name_abbr": "NO",
                        "region": "Europe",
                        "publisher_name": "Funcom",
                        "developer_name": "Funcom"
                    }
                ],
                "current_headcounts": "172",
                "launch_date": "20240930",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "R7 (mid Preproduction Review)",
                        "milestone_node": "",
                        "date": "20220106",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "9",
                        "child": null
                    },
                    {
                        "milestone": "Internal Alpha",
                        "milestone_node": "",
                        "date": "20221205",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "9",
                        "child": null
                    },
                    {
                        "milestone": "Internal Beta 1",
                        "milestone_node": "",
                        "date": "20230629",
                        "date_confirm": "1",
                        "gate_review": "Redo",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "9",
                        "child": null
                    },
                    {
                        "milestone": "Internal R14 Milestone",
                        "milestone_node": "",
                        "date": "20231101",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "9",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20240310",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "9",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta",
                        "milestone_node": "",
                        "date": "20240531",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "9",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "9",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "sdecroix",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 377,
                                    "sort": 0
                                },
                                {
                                    "name": "olegg",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1022",
                                    "id": 447,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "sstellagan",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "4024",
                                    "id": 294,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Joel Bylos",
                                    "position": "CCO and Creative Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 123,
                                    "sort": 0
                                },
                                {
                                    "name": "Scott Junior",
                                    "position": "CTO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 125,
                                    "sort": 0
                                },
                                {
                                    "name": "Shelley Johnson",
                                    "position": "Executive Producer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 129,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "Stian Drageset",
                                    "position": "Funcom CFO and COO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 506,
                                    "sort": 0
                                },
                                {
                                    "name": "Rui Casais",
                                    "position": "Funcom CEO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 507,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "publishing",
                            "list": [
                                {
                                    "name": "Erling Ellingsen",
                                    "position": "Funcom CMO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 505,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Project 070 (B4)",
                "studio_id": "ffbb640e-8df4-4d22-944a-1633b217c2b4",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Platinumgames",
                "project_id": "77",
                "priority": "Key",
                "genre": [
                    "Action"
                ],
                "expanded_genre": "",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Platinum Games",
                "ip": "Existing",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20230926",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "<p>Conditional Pass</p>",
                        "action_items": [
                            {
                                "due_date": "2023-12-31",
                                "user": "PlatinumGames",
                                "action_content": "PlatinumGames to define milestone specifications, and prepare for the mid-gateway for Planning/R&D.",
                                "state": "",
                                "state_comment": ""
                            }
                        ],
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "77",
                        "child": null
                    },
                    {
                        "milestone": "Planning/R＆D",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "77",
                        "child": null
                    },
                    {
                        "milestone": "Prototype",
                        "milestone_node": "",
                        "date": "20250430",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "77",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20251130",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "77",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20260930",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "77",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20270731",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "77",
                        "child": null
                    },
                    {
                        "milestone": "Open Beta",
                        "milestone_node": "",
                        "date": "20271130",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "77",
                        "child": null
                    }
                ]
            },
            "update": [
                {
                    "milestone": "Concept/KOM",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>Bayonetta 1 Remake Pitch</p><p>Not Approved / Canceled</p><p>- Low revenue projection</p><p>- Largely outsourced with little chance for PG staff skill growth</p><p>- Pitched as a marketing tool for Bayonetta 4, but little to no thematic, narrative or visual connections make for a weak case</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.09.24-2023.09.30",
                            "granularity": 2,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-09-29 09:53:48",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><strong>Bayonetta 4 Pitch</strong></p><p>General agreement to proceed</p><p>May provide conditional approval of initial R&D/Concept Phase budget 4.2M USD</p><p>Requirements (Due. Dec. 2023)</p><p>- Detailed list of initial R&D/Concept Phase output</p><p>- Detailed list of risks, mitigations and support requests from Tencent</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.09",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-09-29 09:55:06",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Production</p><ul><li>Planning</li><ul><li>Producer has requested additional time for initial milestone planning deliverables.</li><ul><li>Nov. 10: First deliverable of project specific milestone deliverables</li><li>Nov. 13 - Nov. 24: Round of JVL &lt;&gt; PG feedback</li><li>Nov. 30 - Finalize milestone deliverables</li></ul><li>Schedule up to Dec. 2023 initial milestone funding approval may be found in the B4_FollowUp_Memo &lt; <a href=\"https://drive.weixin.qq.com/s?k=AJEAIQdfAAoZs7AI0l\" target=\"_blank\">https://drive.weixin.qq.com/s?k=AJEAIQdfAAoZs7AI0l</a> &gt;</li><li>Game director took a lenghty leave, his first following Bayonetta 3 and first in several years. As a result game specification planning is behind schedule.</li></ul><li>Impact of Project 053</li><ul><li>Project 053 for Square-Enix (SE) is highest corporate priority.</li><li>053 schedule will have a strong impact on B4 schedule and staffing.</li><ul><li>Barring a cancellation of 053, B4 schedule will most likely be delayed, or focus on a very small initial core team.</li></ul><li>PG to meet with SE the w/o Oct. 30 to decide the fate of the project.</li><ul><li>PG is requesting additional funding and time, but is not the first such request.</li><li>SE may refuse, agree to a smaller amount or cancel the project. The latter seems unlikely, however.</li></ul><li>B4 production schedule will be impacted by 053 schedule</li></ul><li>User Impression Test</li><ul><li>Team to conduct a concept impression test in Nov., and deliver rough results in mid. Dec.</li><li>Impression test results to be used in Dec. initial milestone approval meeting.</li><li>Vivi intermediating discussions with MUR, but cost and time are an issue. Thus suggested local partner PTW as the candidate.</li></ul></ul>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.10",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-10-26 14:41:51",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Planning (Project 070)</p><ul><li>No changes to 2023 schedule below. </li><li>To receive latest round of feedback on milestone definitions on Nov. 23, 2023.</li><li>Milestones</li><ul><li>Nov. 10: First deliverable of project specific milestone deliverables</li><li>Nov. 13 - Nov. 24: Round of JVL &lt;&gt; PG feedback</li><li>Nov. 30 - Finalize milestone deliverables</li></ul><li>Nov. 30, 2023 Deliverables</li><ul><li>Rough GDD/ADD/TDD and Plot Outline</li></ul><li> Early Dec. 2023</li><ul><li>Conduct an domestic impression test utilizing the Game Pitch</li></ul><li>Mid. Dec. 2023 Deliverables</li><ul><li>PG to deliver impression test results report</li></ul><li>~Dec. 31, 2023 Deliverables</li><ul><li>Conduct a review of all aforementioned materials with Steve and Michelle.</li><li>Finalize terms of Bayonetta 4 production contract</li><li>Submit internal application for Bayonetta 4 Production Loan and PG Corporate Loan</li></ul></ul>",
                                "risks": [
                                    {
                                        "Timeline": "Unofficial note from CEO that Dec. milestone may need to be postponed, ostensibly due to the impact of Project 053. No official update from PG at this point."
                                    }
                                ],
                                "workplans": []
                            },
                            "date": "2023.11",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-11-26 21:21:30",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                },
                {
                    "milestone": "Planning/R＆D",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>Planning (Project 070)</p><p>2023 Schedule</p><ul><li>Milestone Definitions</li><ul><li>Dec. 20: Received latest feedback on milestone definitions. Expected to be final round.</li></ul><li>Staffing Plan</li><ul><li>Dec. 20: Vance updated staffing plan data so that current and future versions can be imported to PowerBi, and metrics and changes tracked accordingly.</li></ul><li>Domestic User Impression Test</li><ul><li>Dec. 21: Initial report to be received</li><li>Dec. 26: Final report to be received</li></ul><li>Development Documentation</li><ul><li>~Dec. 31: Updated GDD/ADD/TDD/Plot Outline and Staffing Plan to be received.</li></ul><li>Project Development / Project Loan Contract</li><ul><li>Development contract utilizing pre-existing Project 050 Project Finance contract as base. Reviewed by Tencent Japan finance and legal.</li><li>Development contract was noted to have no reciprocal action from Tencent. Therefore, the Project Loan contract with the action of payment from Tencent in reciprocation for development to serve as the base, with the Development Contract added as an attachment.</li></ul></ul><p>2024 Schedule</p><ul><li>Project Development / Project Loan Contract</li><ul><li>Jan. 12: Current internal deadline for finalizing contract internally.</li></ul><li>~Jan. 31: Present aforementioned materials and plans to Steven and Michelle for approval of initial Concept / R&D Phase. Additionally, seek approval for quarterly project loan.</li></ul>",
                                "risks": [],
                                "workplans": [
                                    {
                                        "comment": "When is the next IEGG check-in point?",
                                        "due_date": "2024-01-31",
                                        "status": 1
                                    }
                                ]
                            },
                            "date": "2023.12",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jinglichen",
                            "update_time": "2023-12-21 22:04:15",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                }
            ],
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "junoshin",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "726",
                                    "id": 542,
                                    "sort": 0
                                },
                                {
                                    "name": "kenimaizumi",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "595",
                                    "id": 543,
                                    "sort": 0
                                },
                                {
                                    "name": "cminami",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "4016",
                                    "id": 544,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Kemuri (PF)",
                "studio_id": "unseen-inc-tokyo",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "UNSEEN Inc.",
                "project_id": "36",
                "priority": "Middle",
                "genre": [
                    "RPG"
                ],
                "expanded_genre": "Online Co-op, Next-gen Social Features",
                "business_model": [
                    "Premium",
                    "Downloadle Content"
                ],
                "image_name": "",
                "developer": "UNSEEN",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "- Kemuri ‘KMR’ is a supernatural, online co-op action game boasting a wild & unpredictable story and a breathtaking anime art style\n- Players take on the role of a Yokai Hunter, teaming up with with NPCs and other online players to confront legions of Yokai, infiltrate shadowy organizations, and uncover the mysteries of a world shrouded in “KEMURI “.",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "ea4c7ec7704246e6bdeec091caed8b946",
                        "unified_id": "ea4c7ec7704246e6bdeec091caed8b946",
                        "name": "Monster Hunter",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/202eb706203f629bef612716d96849b1.jpg",
                        "game_type": 3,
                        "genre": "Action RPG|Misc|Role-Playing",
                        "sub_genre": "",
                        "country_en_name": "Japan",
                        "country_en_name_abbr": "JP",
                        "region": "Japan",
                        "publisher_name": "Capcom",
                        "developer_name": "Capcom"
                    }
                ],
                "current_headcounts": "48",
                "launch_date": "20251130",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Prototype",
                        "milestone_node": "",
                        "date": "20221130",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "36",
                        "child": null
                    },
                    {
                        "milestone": "Pre-production",
                        "milestone_node": "",
                        "date": "20230630",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "36",
                        "child": null
                    },
                    {
                        "milestone": "Beautiful Corner",
                        "milestone_node": "",
                        "date": "20230730",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "36",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "36",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20250430",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "36",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20251130",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "36",
                        "child": null
                    }
                ]
            },
            "update": [
                {
                    "milestone": "Pre-production",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p><strong>HQ Visit - 09/27/23</strong></p><p><u>Overview</u></p><p>Presented Studio Overview, Game Design, Beautiful Corner Video and Hands-on time with 4 Player Co-Op.</p><p><strong>Positives:</strong> Overall very positive impressions on visual and world design. Battle assessed as Devil May Cry-light, which is fairly accurate. </p><p><strong>Negatives:</strong> Co-op gameplay does not foster cooperation at the current stage. Is simply 4 people in the same game acting independently.</p><p><strong>Questions: </strong>Interest shown in the narrative. Not properly shared during presentation, despite a robust narrative already written and planned.</p><p><strong>Finances:</strong> Discussed Unseen request for additional funding.</p><p>・<strong>+6 Months: +</strong><span style=\"color: rgb(0, 0, 0); font-family: Arial;\"><strong>330,075,000</strong></span><strong> JPY</strong></p><p> ・Reclaim 6 Months of cut content</p><p>・<strong>+16 Months: +</strong><span style=\"color: rgb(0, 0, 0); font-family: Arial;\"><strong>548,550,000</strong></span><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\"><strong> JPY</strong></span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・Cumulative: 878,625,000 JPY</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\"> ・Reclaim 6 Months of cut content</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\"> ・10 Months to create Season 2 - Large DLC</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・</span><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\"><strong>+26 Months: +</strong></span><span style=\"color: rgb(0, 0, 0); font-family: Arial;\"><strong>548,550,000 JPY</strong></span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\"> ・Cumulative: </span><span style=\"color: rgb(0, 0, 0); font-family: Arial;\">1,427,175,000 </span><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">JPY</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\"> ・Reclaim 6 Months of cut content</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・10 Months to create Season 2 - Large DLC</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・10 Months to create Season 3 - Large DLC</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\"><strong>Decisions</strong></span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・Steven open to the idea of the +16 Month pattern, but felt the +26 month is too extensive.</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・Discussed funding via a Convertible Loan. Juno to speak further with company founder.</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・JVL, Juno, Yong-yi, Julien B. in agreement that PlayStation is the preferred publisher from a brand perspective. MS is prohibited due to previous bad deals with Tencent partners.</span></p><p><span style=\"color: rgb(31, 31, 31); background-color: rgb(255, 255, 255); font-size: 12px;\">・Unseen floated the idea of self-publishing with Tencent support for QA, Platform, Mktg. support, etc. To reach a final decision on direction by Mar. 2024.</span></p><p><br></p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.09",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-09-28 15:31:06",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Production</p><ul><li>Team continues to focus on Alpha production goals.</li><li>Implementation of Shaman enemy and Bow Hunter player character in progress.</li><li>Environment</li><ul><li>Large reconfiguration of city design to further accentuate verticality.</li><li>Global revisit of previous temporary foliage implementation to accentuate reality.</li><li>Further improvement to environmental world design thanks to the addition of a highly skilled enivorment artist.</li><li>All changes have cleared initial director review</li></ul></ul><p>Publishing</p><ul><li>Continue discussions with SIE, MS and Epic. </li><ul><li>SIE the the highest probability of success, but conditions are TBD. Meeting with SIE rep on Oct. 27, 2023.</li><li>MS has cooled on investment with the success of the ATB merger. New partnership discussions are delayed until 2024.</li><ul><li>Steven and Michelle also noted they would not approve an MS Gamepass deal, due to likely unfavorable conditions.</li></ul><li>Epic is seen as a long shot, and valued for tech support rather than as a potential publisher.</li></ul></ul><p>Promotion</p><ul><li>Work on the TGA 2023 reveal trailer in progress.</li><ul><li>To hand over finalization to agency Petrol in Nov.</li></ul><li>Petrol to deliver promotion roadmap in Nov.</li><li>Landing page and Steam page to be prepared in Unseen.</li><ul><li>EVL Heidi to advise on Steam page creation on Nov. 2 (tentative).</li></ul></ul><p>Investment</p><ul><li>Tencent pushing for further corporate investment.</li><ul><li>To propose conditions to founder on Nov. 1</li><li>Infusion of additional project funding through stock purchase. </li><li>Push for future all option.</li></ul></ul>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.10",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-10-26 14:19:55",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                },
                {
                    "milestone": "Beautiful Corner",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>Production</p><ul><li>Work continues on Alpha and TGA 2023 Trailer</li><ul><li>Nov. 24: TGA 2023 Trailer deadline</li><li>Dec. 15: Close out work for the year</li></ul><li>Current Focus</li><ul><li>Narrative mission implementation. </li><li>Definitive \"This is Kemuri\" narrative game mission / loop</li><li>Large boss implementation</li></ul></ul><p>PR</p><ul><li>Ad Agency - Petrol</li><ul><li>Received Brand Strategy Documents Ver. 1.0 from PR Agency, Petrol</li><ul><li>Content is run of the mill</li><li>Game pillars are not in-line with Unseen's expectations</li><li>Unseen currently collating feedback, and expect to eventually merge their internal plans with Petrols</li></ul><li>Completed second round of Petrol produced key art feedback</li></ul><li>Clothing Line - Ark 8</li><ul><li>Received web page layouts</li></ul><li>Internal</li><ul><li>Continuing to work with internal teams on internal planning</li></ul><li>GVL</li><ul><li>Planning and Marketing Framework workshop with Heidi from EVL in mid. Jan 2024</li></ul></ul><p>Publishing</p><ul><li>MS</li><ul><li>W/O Nov. 13: Conducted first hands-on</li><li>MS strongly praised the game, and team execution capabilities</li><li>Activision Blizzard acquisiton, however, has pushed out all third-party finance discussions to early 2024</li><li>Unseen impressed the desire to reach a timely decision</li></ul><li>SIE</li><ul><li>Dec. 11: Game assessment team to visit</li><li>Dec. 30: An additoinal hands-on scheduled</li></ul><li>HR</li><ul><li>HR</li><ul><li>HR lead to return to work in Dec.</li></ul><li>Animation</li><ul><li>Hired highly skilled animator, formely from film and Ori and the Will of the Wisps, Moon Studios</li><ul><li>Former boss of current Unseen lead animator</li></ul><li>All internal animation slots now filled</li></ul><li>Engineering</li><ul><li>In interview with former Zenimax and current Q-Games rendering engineer</li><ul><li>Strong generalist skill set, and currently interviewing for a Gameplay engineering role</li></ul></ul></ul></ul><p>Investment</p><ul><li>Nov. 30: Juno to speak with founders regarding additional investment with a call option of 80%. 20 to 100%.</li></ul>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.11",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-11-22 16:17:43",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                },
                {
                    "milestone": "Alpha",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>Production</p><ul><li>Production continues on Alpha milestone</li><ul><li>Contractual schedule July 2023 - Sept. 2024</li><li>Dates in flux due to publisher negotiations</li></ul><li>Team had set an internal Vertical Slice-like milestone for Dec. 2023, but creation of the TGA 2023 trailer, and other year-end activities have resulted in a delay of roughly 3 months</li><ul><li>JVL to confirm actual extent of dealy via Hansoft analysis over year-end break.</li></ul><li>New Bow Hunter player character type implemented</li><li>Cinematic camera system implemented</li><li>Battle mechanics for flying enemies implemented</li><li>Large-boss (Demon) proxy model implemented and battle possible</li><li>Additional player character type, Shaman,GDD implementation commenced</li></ul><p><span style=\"color: rgb(68, 68, 68); background-color: rgb(255, 255, 255); font-size: 14px;\">PR</span></p><ul><li><span style=\"color: rgb(68, 68, 68); background-color: rgb(255, 255, 255); font-size: 14px;\">Successfully announced Title at TGA 2023 &lt; </span><a href=\"https://youtu.be/YCAX_eC3rOo?si=oo9Hq5c_hnoA92cn\" target=\"_blank\">https://youtu.be/YCAX_eC3rOo?si=oo9Hq5c_hnoA92cn</a> &gt;</li><ul><li>Largely praised for world design and art style</li><li>Criticized for lack of gameplay, or gameplay-like scenes</li><li>Lack of long-term promotion plan and early timing noted as issues early on, but studio was determined to announce at this time.</li></ul><li>1/29~1/30: GVL Heidi to visit Unseen to conduct Marketing Framework Workshop</li></ul><p>Publishing</p><ul><li>Outreach from MS, Krafton, Microids, IOIGamer following announcement.</li><li>Continue to pursue MS, Sony first-party deal</li><li>Dec. 13: SIE evaluation team visited office for playtest. Results of evaluation pending.</li><li>MS requested a playable build. To share Aug. 8, 2023 pre-pro build in Jan. 2024. Unseen awaiting publisher conditions from MS, which MS claims are delayed due other ongoing projects.</li></ul><p>HR</p><ul><li>HR top returned from health related leave. Currently in rehabilitation.</li><li>Former Moon Studio lead animator to join the team as Unseen lead animator. Current top-tier animator, also from Moon Studio, was former subordinate.</li><li>Environmental artist to leave. Reasons cited pressure of position first time on a small team, first time to finalize a game and adjusting to life in a foreign country.</li></ul>",
                                "risks": [
                                    {
                                        "Financial": "Studio requires additional funding to achieve ideal development plan."
                                    }
                                ],
                                "workplans": []
                            },
                            "date": "2023.12",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-12-24 23:30:38",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Production</p><ul><li>Hansoft (PM Tool)</li><ul><li>Team currently building backlog and planning in Hansoft for 2024</li><li>~Jan. 29: Expect to have latest Sprint input</li></ul><li>Latest Plan (Plan D)</li><ul><li>Concerns from VL that the end of the schedule is unrealistically compressed</li><li>Current Contract Schedule</li><ul><li>Alpha: 2023/7 - 2024/9 (15M / Minus Benchmark = 9M)</li><li>Benchmark (VS) 2023/7 - 2023/12 (6M)</li><li>Beta: 2024/10 - 2025/4 (7M)</li><li>GM Creation: 2025/5 - 2025/7 (3M)</li><li>GM: 2025/8 - 2025/9 (2M)</li><li>Archive: 2025/10 - 2025/11</li></ul><li>Plan D</li><ul><li>Alpha: 2024/1 - 2024/12 (18M / Minus Benchmark = 12M)</li><li>Benchmark 2023/7 - 2023/12 (6M)</li><li>Beta: 2025/1 - 2025/6 (6M)</li><li>GM Creation: 2025/7 - 2025/8 (2M)</li><li>GM: 2025/9 (1M)</li><li>Archive: 2025/10 - 2025/11 (2M)</li></ul><li>Team follow up</li><ul><li>Planning development with the assumption that Alpha will be extended from June 2024 to Dec. 2024.</li><li>Development of Alpha from July 2023 has been proceeding as planned, but additional time is required.</li><li>Unseen development plan is to deliver a relatively bug-free experience at Beta.</li><ul><li>However, Unseen recognized that bugs will remain, and time is required for last-minute localization and sound implementation. Unseen recognized Current Plan D GM Creation period is too short.</li></ul></ul></ul></ul><p>Publishing</p><ul><li>Contract Amendment / Cash Flow</li><ul><li>Current Alpha: Sept. 30, 2024</li><li>Must complete an amendment with an updated schedule prior to the Sept. 30, 2024 deadline</li><li>Amendment will likely require 3 months to complete, therefore updated schedule in lieu of publishing contract should be decided by June. 30, 2024</li><li>Cash Flow is another, and perhaps more pressing concern.</li><ul><li>Payment on a Dec. 31, 2024 Alpha milestone delivery, will likely complete in Mar. 2025 (3M).</li><li>Chisato of TJ Finance to confirm with Unseen how long the current cash reserves will last. </li></ul></ul><li>SIE</li><ul><li>Official response from SIE is expected in Feb.</li><li>Mr. Fujisawa, to escalate via back channels</li></ul><li>MS</li><ul><li>Unseen to share Sept. build with MS this week W/O Jan. 22</li><ul><li>Had planned to send last year, but delayed due to holiday schedule</li></ul></ul><li>Krafton</li><ul><li>First contact and plans for a general introductory meeting</li><li>Krafton has pitched themselves as a publisher willing to support projects of any scope</li><li>Unseen not particularly interested, but will maintain the relationship as a potential secondary plan</li></ul><li>EA</li><ul><li>Direct Email contact via AP, Ms. Fujiwara, following TGA reveal</li><li>As with Krafton, not a strong desire to work with EA, but to investigate the potential further</li></ul></ul><p>PR</p><ul><li>Division of Labor</li><ul><li>AP Ms. Fujiwara and additional PR Staff, Liam, handling majority of PR items</li><ul><li> Ms. Nakamura only involved in final interview stage when required</li></ul></ul><li>Video</li><ul><li>Feb: 3 Videos What is KEMURI, Ask Ikumi 3, Midnight UNSEEN Mr. Nass</li><ul><li>What is Kemuri follows the mold of Bungies What is Marathon Video &lt; <a href=\"https://youtu.be/Zrv0dpztryY?si=yBZTiU8rSkL2Qyxp\" target=\"_blank\">https://youtu.be/Zrv0dpztryY?si=yBZTiU8rSkL2Qyxp</a> &gt;</li></ul></ul><li>Japan Agency of Cultural Affairs</li><ul><li>Feb. 17-22: Special exibition in Tokyo</li><ul><li>Concept Art</li><li>TGA Trailer</li><li>Interview with Ms. Nakamura and famous manga artist</li></ul></ul><li>East Tenessee State University</li><ul><li>To visit studio</li><li>Details pending</li></ul><li>Laurus International School </li><ul><li>Exihibition</li><ul><li>Details to be discussed Jan. 23-24</li></ul></ul><li>Print</li><ul><li>Feb.: Edge Magazine</li><li>Feb.: Vogue Japan</li></ul><li>Online</li><ul><li>IGN</li><ul><li>Featured segment</li><ul><li>Largest interview segment for Unseen thus far</li></ul></ul></ul><li>GDC</li><ul><li>Desire to present regarding development of Kemuri in 2025</li><li>Ken introduced Julien Merceron a long time member of GDC advisory board</li></ul></ul><p>HR</p><ul><li>Producer</li><ul><li>Mr. Shige Tanaka: Discussions put on hold</li></ul><li>Other</li><ul><li>Allen Murray: Founder of Private Division</li><ul><li>Experienced producer, but cannot speak Japanese and has never lived in Japan</li></ul><li>Christofer Lamb: Former Creatures (Pokemon Go), current Shapefarm</li><ul><li>Ken investigating via contacts</li></ul></ul></ul><p>IT</p><ul><li>Tencent OIT (CES) has agreed to fill gap until new IT lead is hired</li><li>AWS/GCP Cost Reduction</li><ul><li>AWS 26%, GCP 30%: Generally reserved for first-party, but to attempt support for Unseen</li></ul></ul><p>Merchandise</p><ul><li>Taking orders for Unseen / KEMURI branded clothing, but quality of initial lot is extremely low</li><li>Ms. Nakamura to work with vendor to improve quality.</li></ul>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2024.01",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2024-01-25 12:46:58",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                },
                {
                    "milestone": "First Playable",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.03",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "wcoveliers",
                            "update_time": "2023-03-15 16:23:37",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                },
                {
                    "milestone": "Prototype",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>Weekly Update ~22/8/29</p><p>1. Hiring</p><p> &nbsp;- Advised on a potential producer hire</p><p> &nbsp; &nbsp; - Resounding NO from Tencent staff</p><p> &nbsp; &nbsp; - The person was undoubtedly maneuvering to take over the company</p><p>2. Planning</p><p> &nbsp;- Currently creating P&L</p><p> &nbsp;- Front BD supporting in securing the Loss element via Unseen’s external accountant</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.09",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2022-09-02 11:51:47",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Development</p><p>1. Showed strong progress at the end of Sept. </p><p>- First time the team collated disparate elements into a mock mission</p><p>- Created a new team dynamic of butting heads and resolving issues in a proactive manner, rather than the previous siloed structure</p><p>- Team to continue working on the mock mission through the end of Oct.</p><p><br></p><p>2. First Milestone, “Prototype” is due Nov. 30</p><p>- To conduct an intensive 2 week “Boot Camp” in preparation for the June 2023 pre-production milestone, as well as scope assessment PR</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.10",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2022-10-26 16:57:19",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p style=\"text-align: start;\"><span><strong>Continuing to work on milestone “Prototype”Prototype</strong></span></p><p style=\"text-align: start;\"><span><strong>- Nov. 30, 2022: Delivery</strong></span></p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.11",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2022-11-17 14:16:44",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Prototype milestone on time. Studio will have a retrospective bootcamp for next year production.</p>",
                                "risks": [
                                    {
                                        "Financial": "Additional project finance ($10M-20M) due to scope change "
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "JVL is conducting milestone review of prototype milestone",
                                        "completed": true,
                                        "due_date": "2022-12-16"
                                    },
                                    {
                                        "comment": "Studio conduct bootcamp of the project issues and next steps for next year production.",
                                        "completed": false,
                                        "due_date": "2022-12-19"
                                    },
                                    {
                                        "comment": "JVL to continue the milestone review of prototype with project team",
                                        "completed": false,
                                        "due_date": "2023-01-16"
                                    },
                                    {
                                        "comment": "JVL to connect with IEGG publishing team and have Tencent team to decide if to publish Kemuri with additional budget. ",
                                        "completed": false,
                                        "due_date": "2023-03-31"
                                    }
                                ]
                            },
                            "date": "2022.12",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "mmiaozhang",
                            "update_time": "2022-12-16 16:02:37",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>The upcoming milestone of pre-production on 6/30 is delaying. Dev team attended GDC to search for the possibiltily of external publisher while negotiating with Level Infinite(Tencent Publishing).</p>",
                                "risks": [
                                    {
                                        "Timeline": "Pre-production on June 30, 2023 will be delay to Aug.31"
                                    },
                                    {
                                        "Production": "Combat system is unsolid"
                                    },
                                    {
                                        "Operational": "Publisher need to be confirmed in time"
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "Continuely support the team to secure the publisher.",
                                        "completed": false,
                                        "due_date": "2023-06-30"
                                    },
                                    {
                                        "comment": "Technical director morita will help on the technical side.",
                                        "completed": false,
                                        "due_date": "2023-06-30"
                                    }
                                ]
                            },
                            "date": "2023.04",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-04-18 17:02:37",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                }
            ],
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "junoshin",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 374,
                                    "sort": 0
                                },
                                {
                                    "name": "jamesvance",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "594",
                                    "id": 538,
                                    "sort": 0
                                },
                                {
                                    "name": "cminami",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "4016",
                                    "id": 539,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Slitterhead:Yakushi (PF) (Minority)",
                "studio_id": "bokehgamestudio",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Bokeh Game Studio",
                "project_id": "30",
                "priority": "Low",
                "genre": [
                    "Action"
                ],
                "expanded_genre": "Horror-adventure",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Bokeh Games Studio",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "ec34d265db8e69c717597202116f75a22",
                        "unified_id": "ec34d265db8e69c717597202116f75a22",
                        "name": "Silent Hill",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/a0f77770c13eaa9c2749de0361eb6d3d.jpg",
                        "game_type": 3,
                        "genre": "Survival|Adventure",
                        "sub_genre": "",
                        "country_en_name": "Japan",
                        "country_en_name_abbr": "JP",
                        "region": "Japan",
                        "publisher_name": "Konami",
                        "developer_name": "KCET"
                    },
                    {
                        "id": "e1067f4f5f6884f00d5512552ddd49b89",
                        "unified_id": "e1067f4f5f6884f00d5512552ddd49b89",
                        "name": "Forbidden Siren 2",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/8a7537dfd5aea970268774ea65530904.jpg",
                        "game_type": 3,
                        "genre": "Action",
                        "sub_genre": "",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "",
                        "developer_name": "SCE Japan Studio"
                    },
                    {
                        "id": "ecb975543cb8615d12627f2725bb01615",
                        "unified_id": "ecb975543cb8615d12627f2725bb01615",
                        "name": "Siren",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/ae1c358e8edfa787cb7b2c11255ef1bc.jpg",
                        "game_type": 3,
                        "genre": "Survival|Adventure",
                        "sub_genre": "",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "SCEA",
                        "developer_name": "SCE Japan Studio"
                    }
                ],
                "current_headcounts": "45+",
                "launch_date": "20240430",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Prototype",
                        "milestone_node": "",
                        "date": "20211031",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "30",
                        "child": null
                    },
                    {
                        "milestone": "Pre-Alpha",
                        "milestone_node": "",
                        "date": "20230430",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "30",
                        "child": null
                    },
                    {
                        "milestone": "Closed Alpha",
                        "milestone_node": "",
                        "date": "20230630",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "30",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20230930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "30",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20231231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "30",
                        "child": null
                    },
                    {
                        "milestone": "Master",
                        "milestone_node": "",
                        "date": "20240331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "30",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20240430",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "30",
                        "child": null
                    }
                ]
            },
            "update": [
                {
                    "milestone": "Closed Alpha",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>As of 2022/7</p><p>- This project is on pre-Alpha phase until the end of 2022</p><p>- Executed several user tests for the demonstration of core game play</p><p>- Team reached out to several publishers for physical publishing (digital published by Bokeh Games itself)</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.08",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2022-09-01 21:59:35",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Upcoming schedule</p><p>- Oct. 2022 Conduct user test of the latest build</p><p>- Nov. 2022 To show JVL user test results and the latest build</p><p> &nbsp; &gt; To prepare presentation pitch(5-10pages) for Steven’s visit</p><p>- Q1 2023: Present and show the latest build to Steven Ma upon his visit to Japan</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.10",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2022-10-26 16:59:50",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><span style=\"background-color: rgb(255, 255, 255); font-size: 14px;\">Oct. 17: Conducted user test of the latest build (2nd test)</span></p><p><span style=\"background-color: rgb(255, 255, 255); font-size: 14px;\">Nov. 16: Showed JVL the latest build (middle pre-Alpha)</span></p><p>Confirmed self-publish digital version</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.11",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "mmiaozhang",
                            "update_time": "2022-12-16 15:28:37",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Successfully transitioned to Unreal Engine 5. Pre-Alpha production is on Track. Had 2 user tests for improving gameplay this year.</p>",
                                "risks": [
                                    {
                                        "Timeline": "Publishers need to be confirmed in time"
                                    },
                                    {
                                        "Production": "Graphic and Action function need to be improved"
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "JVL conduct on-site test play of build, confirm production progress and discuss development and publisher needs. ",
                                        "completed": false,
                                        "due_date": "2023-01-31"
                                    },
                                    {
                                        "comment": "In progress of partnering with Marvelous for physical version publishing in Japan. ",
                                        "completed": false,
                                        "due_date": "2023-01-31"
                                    },
                                    {
                                        "comment": "In progress of partnering with FireShine for physical version publishing  outside of Japan.",
                                        "completed": false,
                                        "due_date": "2023-01-31"
                                    }
                                ]
                            },
                            "date": "2022.12",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "wcoveliers",
                            "update_time": "2023-03-15 16:23:37",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Team is preparing for the upcoming milestone \"Pre-Alpha\" on Apr. 2023 for Tencent's internal review. $5,000,000 will be funded upon the approval of Tencent.</p>",
                                "risks": [
                                    {
                                        "Timeline": "The announcement and release plan/date need to be confirmed."
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "JVL will review the \"Pre-Alpha\" milestone promptly after submited.",
                                        "completed": false,
                                        "due_date": "2023-05-19"
                                    }
                                ]
                            },
                            "date": "2023.04",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-04-18 17:08:19",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                }
            ],
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "junoshin",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 375,
                                    "sort": 0
                                },
                                {
                                    "name": "weivivizhou",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "10825",
                                    "id": 445,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Kestrel",
                "studio_id": "579787ca-907c-f5ae-9567-7b99a6c39039",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Remedy Entertainment",
                "project_id": "82",
                "priority": "Middle",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "",
                "business_model": [
                    "Premium",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Tencent",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "15",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "82",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "gramxu",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "290",
                                    "id": 560,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Project Edo (PF)",
                "studio_id": "87629a56-0325-11ec-bc3d-525400d129b6",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Soleil Ltd.",
                "project_id": "41",
                "priority": "Key",
                "genre": [
                    "Action"
                ],
                "expanded_genre": "Multiplayer Action PvPvE",
                "business_model": [
                    "Premium",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "Soleil",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "Project Edo is a PvPvE close-quarters, multiplayer, online action game leveraging the precision combat and Ninja genre expertise of the Soleil team.",
                "ambition": "The project will prove an integral step in growing Soleil as a company by establishing a successful break from work-for-hire operations, and engaging the GaaS business model.",
                "similar_games": [
                    {
                        "id": "ea3a8c0f4c15ae15a875184e6fa700dbf",
                        "unified_id": "ea3a8c0f4c15ae15a875184e6fa700dbf",
                        "name": "Naraka: Bladepoint",
                        "cover": "https://ogdb-cdn.intlgame.cn/pic_by_handle/naraka bladepoint.jpg",
                        "game_type": 2,
                        "genre": "Action Games|Battle Royale|Hack and slash/Beat 'em up|3D Beat-'Em-Up|Action|MMO|Action & adventure|Fighting|Azione e avventura|Action et aventure|アクション & アドベンチャー|Acción y aventura|Actie en avontuur",
                        "sub_genre": "",
                        "country_en_name": "China",
                        "country_en_name_abbr": "CN",
                        "region": "China(Mainland)",
                        "publisher_name": "NetEase",
                        "developer_name": "24 Entertainment"
                    },
                    {
                        "id": "ecbfe9d3f48a2ca00a275c29e94790bc9",
                        "unified_id": "ecbfe9d3f48a2ca00a275c29e94790bc9",
                        "name": "Escape From Tarkov",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/225769c0a7baf9ebffda0f3dda7d162f.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Simulator|Role-playing (RPG)|Tactical",
                        "sub_genre": "",
                        "country_en_name": "United Kingdom",
                        "country_en_name_abbr": "GB",
                        "region": "Europe",
                        "publisher_name": "Battlestate Games",
                        "developer_name": "Battlestate Games"
                    }
                ],
                "current_headcounts": "90+",
                "launch_date": "20250930",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20221231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": null
                    },
                    {
                        "milestone": "Prototype",
                        "milestone_node": "",
                        "date": "20230331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20240229",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20240430",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": [
                            {
                                "milestone": "-Alpha #2 (Alpha with Teamplay first playable)",
                                "milestone_node": "",
                                "date": "20240731",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "41"
                            },
                            {
                                "milestone": "-Alpha #3 (Alpha with Teamplay VS)",
                                "milestone_node": "",
                                "date": "20241030",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "41"
                            }
                        ]
                    },
                    {
                        "milestone": "Soft launch",
                        "milestone_node": "",
                        "date": "20241231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": null
                    },
                    {
                        "milestone": "Gold Master Candidate",
                        "milestone_node": "",
                        "date": "20250331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20250530",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20250930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "41",
                        "child": null
                    }
                ]
            },
            "update": [
                {
                    "milestone": "Prototype",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.05",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "wcoveliers",
                            "update_time": "2023-05-19 18:43:01",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Development of Vertical Slice proceeding as planned. Plan to complete the VS Phase 1 build with action gameplay on Sep. for IEGG's review to decide publishing.</p>",
                                "risks": [
                                    {
                                        "Technological": "Art target as Arcane is set too high that the team hasn't proven the feasibility."
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "Continue to support the team on VS build and presentation to IEGG.",
                                        "completed": false,
                                        "due_date": "2023-09-30"
                                    }
                                ]
                            },
                            "date": "2023.06",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-06-08 15:42:20",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><strong>Overall:</strong> Development of Vertical Slice proceeding as planned. Plan to complete the VS Phase 1 build with improved action gameplay prior to W/O Sep. 25 IEGG publishing review. Steven, Michelle, and potentially Martin may visit Japan to review, in addition to Kiai and the Japan IEGG publishing team.</p><p><strong>Art: </strong>June 27 - Latest environment art direction and guidelines as well as character creation guidelines shared with JVL. July 6, JVL advisor, Mr. Sasaki, provided feedback and guidance on HLOD, Memory Management, Foreground, Mid-ground and Background Consistency, Guideline Direction and Color Space.</p><p><strong>PM: </strong>June task completion ratio of 60%. Incomplete tasks generally centered on out game elements. Overflow carried over to July with plan to recover lost time mapped out in Hansoft. In-game elements remain on track. Venture Lab team and Juno to visit Soleil office on July 13 for hands-on with the latest build.</p><p>User Tests: Sept. - Test the Aug. VS Phase 1 build Battle and Stealth elements prior to IEGG publishing team visit. Aug-Oct. VS Phase 2 test dependent on schedule and progress. Dec. VS test currently planned. Discussions regarding the use of MUR ongoing.</p><p><strong>Corporate: </strong>JVL working with finance to complete 5-year plan, and verify validity of subsequent Project Edo sales/revenue projections.</p>",
                                "risks": [
                                    {
                                        "Production": "Art continues to remain a concern. Team is proceeding via the following schedule to address the issue.\nJune: Determine the art overall fidelity\nJuly: Subset of models featuring final art quality\nAug.: Additional polish and finalization of overall quality target\n"
                                    }
                                ],
                                "workplans": []
                            },
                            "date": "2023.07",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-07-11 13:04:34",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><strong>Overall</strong></p><p>Soleil continues to focus on Vertical Slice Phase 1 production, to be complete at the end of Aug. Also making preparations for IEGG Publishing team / Steven Ma and company presentation in Sept. - Oct.</p><p><strong>Production</strong></p><p>Production updates, team revisited aspects of game design based on GRC and Databrain feedback, and have determined to switch from an open weapon based system to a system in which each weapon type is restricted to a particular hero. Additionally, the number of heroes available at launch will be reduced from 13 to 7, but the rate at which the heroes are released post-launch will be in line with Naraka and other competitive titles at 1 new hero character every two months. Note that this may potentially reduce to once every three months after year two.</p><p>Additional changes include the reduction of total launch timing areas from 4 to 6, with the idea to rotate open areas at regular intervals as a means to foster variety of play. The total number of players to be reduced from 100 to 60.</p><p><strong>User Test / MUR</strong></p><p>Vivi had secured an MUR testing slot for Project G.G. on Sept. 2-3, but due to the project cancellation this slot has been shifted to Project Edo.</p><p>■Schedule</p><p>・W/O Aug. 1: </p><p>Soleil confirms the desired User Test content and planning</p><p>・W/O Aug. 7: </p><p>MUR team to produce testing plan first draft based on Soleil proposal. </p><p>Soleil to share latest build &gt; may prove difficult to make server preparations by this time. To confirm with Soleil during Aug. 10 on-site meeting.</p><p>・W/O Aug. 14:</p><p>MUR and Soleil to reach an agreement on final user test plan.</p><p>・W/O Aug. 28: </p><p>Soleil to provide latest test build.</p><p>・Sept. 1: </p><p>Test on-site preparations, build installation, etc. Soleil PM, Mr. Inamori, likely to attend to ensure smooth operation and deal with any potential issues.</p><p>・Sept. 2-3</p><p>Conduct user test. At Japan testing vendor site.</p><p>・Sept. 11</p><p>・MUR provides final user test report.</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.08",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-08-03 16:20:00",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p style=\"text-align: start;\"><strong>Overall</strong></p><p style=\"text-align: start;\">Soleil has entered Vertical Slice Phase 2 of 3 as of Aug. 1. Tracking 62.3% completion rate of August tasks in Hansoft as of Aug. 24.</p><p style=\"text-align: start;\">Vertical Slice Phase 1: April 1 - July 31.</p><p style=\"text-align: start;\">Vertical Slice Phase 2: Aug. 1 - Oct. 31.</p><p style=\"text-align: start;\">Vertical Slice Phase 3: Nov. 1 - Dec. 31</p><p style=\"text-align: start;\">A user test facilitated by MUR to be conducted on Sept. 2-3.</p><p style=\"text-align: start;\">IEGG Publishing team will visit Soleil on Sept. 14 in preparation for Steven Ma and company presentation on Sept. 27. Tencent Japan conducting ongoing discussions with Publishing Team, and collaboration on forecast. Ideally, would like to meet a publishing decision to be determined shortly after meeting.</p><p style=\"text-align: start;\"><strong>Production</strong></p><p style=\"text-align: start;\">No significant updates. </p><p style=\"text-align: start;\"><strong>User Test / MUR</strong></p><p style=\"text-align: start;\">■Remaining Schedule</p><p style=\"text-align: start;\">・Sept. 1:</p><p style=\"text-align: start;\">Test on-site preparations, build installation, etc. Soleil PM, team to attend to ensure smooth operation and deal with any potential issues.</p><p style=\"text-align: start;\">・Sept. 2-3</p><p style=\"text-align: start;\">Conduct user test. At Japan testing vendor site.</p><p style=\"text-align: start;\">・~Sept. 15</p><p style=\"text-align: start;\">・MUR to provide final user test report.</p>",
                                "risks": [
                                    {
                                        "Production": "Art competency remains a risk. A retail quality section to be assessed during Sept. 2-3 user test."
                                    }
                                ],
                                "workplans": []
                            },
                            "date": "2023.09",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-08-31 15:35:32",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><strong>HQ visit to studio: Sep. 27, 2023</strong></p><p><u>Overview</u></p><p><strong>Content:</strong> Game Overview, Development Progress Update, Hands-on, Financing and Publishing Discussion</p><p><strong>Positives:</strong> Steven and IEGG Publishing team positive on the game design and production plan.</p><p><strong>Negatives:</strong> See current lack of team-based mode and social aspect as a risk. Though team was one step ahead, and had prepared a team based mode content production and budget plan in advance.</p><p><u>Decisions</u></p><p><strong>Publishing</strong></p><p>・Steven, Michelle and Kiai agreed to publish Project EDO under the Level Infinite brand.</p><p><strong>Finances</strong></p><p>・+11M USD for remaining Post-VS to Master Production</p><p>・+12M USD for Year 1 GaaS content production</p><p>・+7-10M USD for Team Mode / Social functionality production to be released +6 months after launch.</p><p>*Note: Total ~42M USD, including current 10M USD project finance, to be converted to Project Loan held by Soleil </p><p><strong>Conditions</strong></p><p>・Must hit 3-4M units in first year to be competitive with current leader, Naraka: Bladepoint. &nbsp;</p><p>・Soleil to work with IEGG Publishing to formulate a competitive, and feasible content production plan.</p><p>・Soleil to work with IEGG to conduct user research to ensure that the Japanese theme content is presented in such a way so that it can be readily understood and digested by western audience.</p><p>・IEGG to work with Soleil to ensure that infrastructure and operations are hardened and tested for release, and that no historical errors are repeated. e.g. Launch connection errors, etc. </p><p><strong>Other</strong></p><p>・IEGG revenue share discussions to continue between IEGG Publishing </p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.10",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jamesvance",
                            "update_time": "2023-09-28 15:19:48",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><strong>Meetings</strong></p><p>Oct 18: Intro meeting with DD from PPS GaaS team</p><p>- Mainly information gathering about the game design</p><p>- DD will follow on GaaS topic and support on the VS gate review in terms of GaaS. Next on-site meeting is Nov.16 </p><p>Oct.19: Bi-Weelky production meeting</p><p>- Production is generally proceeding well. Applied UE 5.2 in early Oct.</p><p>- As of 10/19, overall task progression is 42%, which is mostly as planned</p><p>Oct. 30: IEGG Publishing team to visit Solel, to discuss on pricing and future production plan </p><p><strong>Vertical Slice </strong></p><p>VS build scope as below:</p><p>-Sector(map): 4</p><p>-Weapon: 4</p><p>-Mission: 20~30 (including PvE)</p><p>-Training Hall: Implement</p><p>-Trader(shop): Implement</p><p>-Hideout (build-up facilities): No implementation, planning in Alpha ver.</p><p>-Art100% Final quality as release product (Both characters and environment)</p><p>JVL concern on the limitation of art delivery due to the absence of AD and TA</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2023.11",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-10-26 14:57:40",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><strong>Production</strong></p><p>Team continues to focus on Vertical Slice production goals.</p><p><strong>Meetings</strong></p><p>Nov. 16 Bi-weekly meeting and hands-on playtest</p><p> - 46.9% completion of Nov. tasks as of 11/16. </p><p> - Generally on schedule, while outgame production is going behind.</p><p><strong>Gate Review</strong></p><p>Nov. 21 Had a introductory meeting with IEGG publishing team (Kiai's team)</p><p> - Decided to conduct GR after JVL's internal review (need to align with Juno)</p><p> - Estimation Date: </p><p> &nbsp; 12/29 Delivery of VS deliverables</p><p> &nbsp; &nbsp;Mid-January- JVL's review</p><p> &nbsp; &nbsp;Feburary~ GR</p><p><strong>User Test</strong></p><p>IEGG pub team proposed a qualitative user test in North American market in next Jan.. Base on the VS delivered on Dec. (6 people ×4 teams=24 in total)</p>",
                                "risks": [
                                    {
                                        "Production": "Overall graphic was unpolished"
                                    },
                                    {
                                        "Production": "Action feel and playability is not balanced"
                                    }
                                ],
                                "workplans": []
                            },
                            "date": "2023.11.19-2023.11.25",
                            "granularity": 2,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-11-22 15:13:51",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Dec 14: JVL held meeting with Soleil on the preparation of VS Gate Review and release schedule</p><p><br></p><ul><li>Challenges to VS Gate Review</li></ul><p> 1) Graphic: Not proven as “final quality”</p><p> - Especially environment art is significantly below average</p><p> - Team claimed they were too focused on gameplay/frame rate thus deprioritized the environment art, but JVL suggested to show final quality in certain area despite of the quality to prove team’s ability</p><p> - JVL technical director Tom worked on-site from Dec. 18 &nbsp;to identify potential performance issues, particularly those that may be holding back art pipeline</p><p><br></p><p> 2) Gameplay: Need to prove “fun” in 60 players environment</p><p> - Game is designed for 60 PvPvE but not yet be tested by Tencent in that context (Planning on 1/10 after VS's delivery) &gt; Tom did and gave positive feedback</p><p><br></p><p> 3) Other</p><p> - Major change in production direction: Team play mode</p><p> - EDO was designed and developed for solo play originally, but after Steven/MM's visits and the discussion with IEGG Publishing team in Dec, team decided to apply team play mode as the main mode on the release to reach the sales target</p><p> - As this is not applying to Dec. VS build (technically not \"final quality\"), we need to align with internal team on GR</p><p><br></p><ul><li>Release Schedule</li></ul><p>1) Soleil made a unrealistic production plan in order to meet Tencent’s original demand on June 2025 release</p><p> - JVL suggested Soleil to adjust the plan into more practical one to avoid future conflicts with IEGG, and guarantee the QA</p><p> - After discussion and alignment with IEGG, Soleil decide to revise the production plan with release on Sep. 2025 (3 month delay)</p><p><br></p><p>2) Updated schedule (WIP)</p><ul><li> 2023 Dec.: VS (Soloplay VS)</li><li>2024 Apr.: Alpha#1 (Teamplay prototype)</li><li>2024 July: Alpha#2 (Alpha with Teamplay first playable)</li><li>2024 Oct.: Alpha#3 (Alpha with Teamplay VS)</li><li>2025 Feb: Beta#1 (Beta)</li><li>2025 May: Beta#2 (Beta with QA)</li><li>2025 Aug: GM</li><li>2025 Sep.: Launch</li></ul><p><br></p><p>3) Suggest GR schedule</p><ul><li>VS GR: 2024 Q1(Mar.-Apr.) Soloplay VS</li></ul><p> - Need to align the goal with assessment team </p><p> - Need to pass JVL review on Jan. 2024</p><ul><li>Alpha GR: &nbsp;2024 Oct. Teamplay VS+Alpha </li></ul>",
                                "risks": [
                                    {
                                        "Production": "Enviornment graphic does not meet the final quality"
                                    },
                                    {
                                        "Production": "PM tool (Hansoft) is not used efficiently"
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": " JVL technical director Tom worked on-site from Dec. 18  to identify potential performance issues, particularly those that may be holding back art pipeline",
                                        "due_date": "2023-12-22",
                                        "status": 2
                                    },
                                    {
                                        "comment": "Align with DPS and internal team on GR schedule",
                                        "due_date": "2023-12-29",
                                        "status": 2
                                    }
                                ]
                            },
                            "date": "2023.12",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-12-21 16:22:12",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                },
                {
                    "milestone": "Vertical Slice",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>Soleil (Project Edo: Vertical Slice - Retry) &gt; In Progress</p><ul><li>Production</li><ul><li>Dec 28: VS milestone Delivery</li><ul><li>JVL Review Result: Failed/Retry</li><li>Jan. 19：Final report of JVL shared to Soleil</li><li>Jan. 20: Meeting with Soleil to share details of VS review and expected actions</li></ul></ul><li>Upcoming schedule</li><ul><li>Jan 30. : Soleil presents new art benchmark goals and resubmission deadlines</li><ul><li>New art benchmark title and logic</li><li>1st Draft updated production plan (including release delays, if any)</li><li>1st Draft updated Man-months</li></ul><li>~Feb. 6: Tencent and Soleil reach an agreement on goals and plan</li><li>Feb. 6: Meeting with finance team to get acquire project finance approval</li></ul><li>Gate review</li><ul><li>Date TBD (depending on the results of resubmission)</li></ul><li>Publishing Contract</li><ul><li>~Dec. 31, 2023: IEGG & Soleil reached agreement on contract terms</li><li>IEGG signature contingent on Gate review/OR(oversea review)</li></ul><li>User Test</li><li> Jan. 25 - 26 (PST): IEGG Pub team leading the implementation of user test in North American market via the vendor Interpret</li><ul><li>Version: VS</li><li>Attendees</li><ul><li>IEGG Publishing: Dale, Icy, Claire</li><li>JVL: Vivi, Ken</li><li>Soleil: Hiragami (Director), Inamori (PM) </li></ul><li>Detail plan and discussion guide delivered and agreed upon</li></ul></ul>",
                                "risks": [
                                    {
                                        "Production": "VS failed due to failure of art to meet pre-established design and quality benchmarks."
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "Update GPOP milestone and VS internal review results",
                                        "due_date": "2024-02-02",
                                        "status": 1
                                    }
                                ]
                            },
                            "date": "2024.01",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "jinglichen",
                            "update_time": "2024-01-25 15:33:28",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                },
                {
                    "milestone": "Pitch",
                    "update_list": [
                        {
                            "content": {
                                "key_updates": "<p>Weekly Updates ~2022/8/29</p><p><br></p><p>1. Development</p><p>Sept. 1: Weekly Meeting</p><p>Aug. 31: To receive non-contract deliverables </p><p><br></p><p>2. Project Managment</p><p>Aug. 17: Mr. Sasaki, advisor conducted a presentation on Hansoft set-up and usage</p><p>Aug. 31: Mr. Sasaki to conduct part two of Hansoft presentation, diving into advanced features and operations</p><p><br></p><p>3. Art</p><p>Aug. 19: At Mr. Sasaki request Soleil provided a list of art and art-tech related development inquiries (Questions) </p><p>Aug. 22: Mr. Sasaki provided responses to the aforementioned inquiries (Responses)</p><p>Aug. 22: Mr. Sasaki sent a questionnaire regarding art staffing structure and personnel and art tools to Soleil (Questions)</p><p>Aug. 25: Received response from Soleil (Responses)</p><p> &nbsp; - Impressions are that Soleil core staff are severely lacking, particularly considering the number of production lines</p><p> &nbsp; - As JVL is aware, Soleil’s reliance on outsourcing and contractors remains evident</p><p> &nbsp; - General experience with most art development tools, with the exception of Houdini, which is not unusual as experts in Houdini are rare</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.09",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2022-09-01 20:58:28",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Development</p><p>- Oct. 3 hands-on with teamBuild demonstrated core concepts of close-quarters combat, animal transformation and online multiplayer and Tarkov-like scavenging</p><p> &nbsp;- Currently focused on demo build for Steven Ma’s visit in Q1 2023</p><p> &nbsp;- Monetization Team is admittedly inexperienced and along with the CEO, have asked for Tencent support</p><p>- Oct. 14: James and Ken had first meeting with Chris W. </p><p> &nbsp;-To report findings to Soleil team and work together to develop initial monetization plan </p><p> &nbsp;-Chris’ experience on Escape from Tarkov may also prove invaluable to the Edo team</p>",
                                "risks": [],
                                "workplans": []
                            },
                            "date": "2022.10",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2022-10-26 16:34:36",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><br></p><p>Continuing to work on the next milestone “Pre-Production” and present build to Steven Ma in Q1 2023</p><p><br></p>",
                                "risks": [],
                                "workplans": [
                                    {
                                        "comment": "Pre-Production Delivery",
                                        "completed": false,
                                        "due_date": "2022-12-30"
                                    },
                                    {
                                        "comment": "Present build to Steven Ma in Q1 2023",
                                        "completed": false,
                                        "due_date": "2023-03-31"
                                    }
                                ]
                            },
                            "date": "2022.11",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "mmiaozhang",
                            "update_time": "2022-12-15 08:52:47",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Current milestone expected to complete ahead of schedule. Work on a special build for Steven's visit on Q1 2023.</p>",
                                "risks": [
                                    {
                                        "Production": "Graphics can be more competitive and appealing but with technical difficulty"
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "JVL to review the Pre-production milestone deliverables",
                                        "completed": false,
                                        "due_date": "2023-01-16"
                                    },
                                    {
                                        "comment": "JVL to conduct on-site special build test ",
                                        "completed": false,
                                        "due_date": "2023-01-31"
                                    },
                                    {
                                        "comment": "JVL consultant Sasaki and technical director Tom to advice on graphics team",
                                        "completed": false,
                                        "due_date": "2023-03-31"
                                    }
                                ]
                            },
                            "date": "2022.12",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "mmiaozhang",
                            "update_time": "2022-12-16 15:56:39",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Team has delivered \"Pre-production\" milestone and JVL/Juno decided to give \"PASS\"</p><p>Team continuing to prepare/brush up the playable build for Steven’s visit.</p>",
                                "risks": [
                                    {
                                        "Production": "Graphics can be more competitive and appealing but with technical difficulty"
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "JVL consultant Sasaki and technical director Tom to advice on graphics team",
                                        "completed": true,
                                        "due_date": "2023-03-31"
                                    }
                                ]
                            },
                            "date": "2023.01",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-04-18 15:51:18",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p><span style=\"color: rgb(0, 0, 0); background-color: rgb(255, 255, 255); font-size: 14px;\">Team has delivered \"Prototype\" milestone and JVL/Juno decided to give \"PASS\". On </span>4/7, team presented to Steven and HQ team and recieved overall highly positive impressions of hands-on gameplay/design.</p>",
                                "risks": [
                                    {
                                        "Production": "Graphics can be more competitive and appealing but with technical difficulty"
                                    },
                                    {
                                        "Production": "Maps/level designs is too complicated and need to be optimized"
                                    },
                                    {
                                        "Operational": "Need to secure addintional funding after Verticle slice"
                                    },
                                    {
                                        "Operational": "Need to secure the publisher"
                                    },
                                    {
                                        "Production": "Action mechanics, while sound, require additional polish"
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "Milestone review on the \"Prototype\" milestone by JVL, and give constructive feedback on the risks.",
                                        "completed": true,
                                        "due_date": "2023-03-31"
                                    },
                                    {
                                        "comment": "Continuely supporting the dev team on production and user testing",
                                        "completed": false,
                                        "due_date": "2023-12-31"
                                    }
                                ]
                            },
                            "date": "2023.04",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-04-18 15:53:23",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        },
                        {
                            "content": {
                                "key_updates": "<p>Development of Vertical Slice proceeding as planned. Steven approved of IEGG publishing, but IEGG publishing team (Kiai) is verifying and plan to make a decision in Late Aug.</p>",
                                "risks": [
                                    {
                                        "Production": "Graphics can be more competitive and appealing but with technical difficulty"
                                    },
                                    {
                                        "Production": "Maps/level designs is too complicated and need to be optimized"
                                    },
                                    {
                                        "Operational": "Need to secure addintional funding after Verticle slice"
                                    },
                                    {
                                        "Operational": "Need to secure the publisher"
                                    },
                                    {
                                        "Production": "Action mechanics, while sound, require additional polish"
                                    }
                                ],
                                "workplans": [
                                    {
                                        "comment": "JVL to support he team to secure publisher (IEGG as the first priority) ",
                                        "completed": false,
                                        "due_date": "2023-08-31"
                                    },
                                    {
                                        "comment": "Continuely supporting the dev team on production and user testing",
                                        "completed": false,
                                        "due_date": "2023-12-31"
                                    }
                                ]
                            },
                            "date": "2023.05",
                            "granularity": 1,
                            "progress": 0,
                            "update_user": "weivivizhou",
                            "update_time": "2023-05-11 16:36:40",
                            "view_update": false,
                            "milestone": "",
                            "milestone_id": "",
                            "id": 0
                        }
                    ]
                }
            ],
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "junoshin",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 373,
                                    "sort": 0
                                },
                                {
                                    "name": "weivivizhou",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "10825",
                                    "id": 540,
                                    "sort": 0
                                },
                                {
                                    "name": "cminami",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "4016",
                                    "id": 541,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Crysis 4",
                "studio_id": "a5046e37-9d89-fd14-069a-1df7924d66ed",
                "platform": [
                    "pc"
                ],
                "studio": "Crytek",
                "project_id": "45",
                "priority": "Low",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "Tactical FPS",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Unknown",
                "ip": "Existing",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20261130",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pre-pitch",
                        "milestone_node": "",
                        "date": "20220502",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "45",
                        "child": null
                    },
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20221230",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "45",
                        "child": null
                    },
                    {
                        "milestone": "Prototype",
                        "milestone_node": "",
                        "date": "20231130",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "45",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20240831",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "45",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20251130",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "45",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20260930",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "45",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20261130",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "45",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "gramxu",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 359,
                                    "sort": 0
                                },
                                {
                                    "name": "alicezqlu",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "19182",
                                    "id": 561,
                                    "sort": 0
                                },
                                {
                                    "name": "samaelwang",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "11580",
                                    "id": 562,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Exterminauts (PF)",
                "studio_id": "f0b19907-a6fe-5d0d-c44a-9e3db9ddb256",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Behaviour Interactive",
                "project_id": "31",
                "priority": "Middle",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "Semi-openworld + RPG looter shooter",
                "business_model": [
                    "P+MTX"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "e11000000282",
                        "unified_id": "e11000000282",
                        "name": "Warframe",
                        "cover": "https://ogdb-cdn.intlgame.cn/pic_by_handle/Warframe.jpeg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Role Playing|Misc|Third Person Shooter|Free to Play Games|Role-playing (RPG)|Action|Third-Person Shooter",
                        "sub_genre": "",
                        "country_en_name": "Canada",
                        "country_en_name_abbr": "CA",
                        "region": "North America",
                        "publisher_name": "Digital Extremes",
                        "developer_name": "Digital Extremes"
                    },
                    {
                        "id": "e11000000288",
                        "unified_id": "e11000000288",
                        "name": "Destiny 2",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/e7a0f4301d1f88cbf9f79296cbca7c9a.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Action Games|First-Person Shooter|RPG|Role-playing (RPG)|Action|MMO|Action & adventure|FPS|Azione e avventura|Action et aventure|アクション & アドベンチャー|Acción y aventura",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Bungie",
                        "developer_name": "Bungie"
                    }
                ],
                "current_headcounts": "",
                "launch_date": "20250531",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Milestone 1 – End of Pre-Production",
                        "milestone_node": "",
                        "date": "20220510",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    },
                    {
                        "milestone": "Milestone 2 - Prototype",
                        "milestone_node": "",
                        "date": "20220810",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    },
                    {
                        "milestone": "Milestone 2.5 - First playable I",
                        "milestone_node": "",
                        "date": "20221110",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    },
                    {
                        "milestone": "Milestone 3 - First playable II",
                        "milestone_node": "",
                        "date": "20230127",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    },
                    {
                        "milestone": "Milestone 4 – Start of Production",
                        "milestone_node": "",
                        "date": "20230210",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    },
                    {
                        "milestone": "Milestone 5 - Mid-Vertical Slice",
                        "milestone_node": "",
                        "date": "20230717",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    },
                    {
                        "milestone": "MS6 - Vertical Slice",
                        "milestone_node": "",
                        "date": "20240222",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20250531",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "31",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "gramxu",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 413,
                                    "sort": 0
                                },
                                {
                                    "name": "Leonkang",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "12304",
                                    "id": 484,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "-",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "",
                                    "id": 331,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "T3-Dungeon Tour",
                "studio_id": "cc69ad75-2fa2-a3f5-5be2-40fa618f141b",
                "platform": [
                    "mobile",
                    "pc",
                    "console"
                ],
                "studio": "Tequila Works",
                "project_id": "79",
                "priority": "Middle",
                "genre": [
                    "Casual"
                ],
                "expanded_genre": "",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Tequila Works",
                "publisher": "Tequila Works",
                "ip": "New",
                "tagline": "4-player; Co-op; Party game; Casual.\nCooperative multiplayer casual party game, Overcooked, Goblins vs.Tourists",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20251231",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20231103",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "79",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20240731",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "79",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta",
                        "milestone_node": "",
                        "date": "20250331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "79",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20250630",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "79",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20251231",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "79",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "748",
                                    "id": 545,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Hydra",
                "studio_id": "80b39a2b-6b55-4066-868b-029bbd2bc8b1",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Fatshark AB",
                "project_id": "7",
                "priority": "",
                "genre": [
                    "TBC"
                ],
                "expanded_genre": "TBD",
                "business_model": [
                    "TBD"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "Extraction, Fantasy Nordic, 1P melee combat",
                "ambition": "TBC",
                "similar_games": null,
                "current_headcounts": "12",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20230331",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    },
                    {
                        "milestone": "Concept",
                        "milestone_node": "",
                        "date": "20230830",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    },
                    {
                        "milestone": "Concept review",
                        "milestone_node": "",
                        "date": "20240229",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    },
                    {
                        "milestone": "Prototype",
                        "milestone_node": "",
                        "date": "20240430",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    },
                    {
                        "milestone": "Pre-production",
                        "milestone_node": "",
                        "date": "20240731",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20250531",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20251031",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20260331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "7",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "gramxu",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 357,
                                    "sort": 0
                                },
                                {
                                    "name": "Victorshen",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "208",
                                    "id": 424,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Nightingale",
                "studio_id": "ca3001c7-2af3-4498-8296-a33ff83da153",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Inflexion Games",
                "project_id": "46",
                "priority": "Top 5",
                "genre": [
                    "RPG/SOC"
                ],
                "expanded_genre": "Co-op PvE",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Inflexion Games",
                "publisher": "Inflexion Games",
                "ip": "New",
                "tagline": "Prepare for a journey of adventure, danger, and discovery - as you search for a way back to the last haven of humanity, Nightingale.",
                "ambition": "FPS Co-op PvE set in Victorian Gaslamp Fantasy",
                "similar_games": null,
                "current_headcounts": "106",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20220831",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "46",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20221124",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "46",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20231222",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "46",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20240222",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "46",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "francoischa",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "14648",
                                    "id": 415,
                                    "sort": 0
                                },
                                {
                                    "name": "faudet",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "4020",
                                    "id": 486,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 333,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Aaryn Flynn",
                                    "position": "CEO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 134,
                                    "sort": 0
                                },
                                {
                                    "name": "Neil Thompson",
                                    "position": "Art & Audio Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 136,
                                    "sort": 0
                                },
                                {
                                    "name": "Leah Summers",
                                    "position": "Senior Producer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 523,
                                    "sort": 0
                                },
                                {
                                    "name": "Scott Nye",
                                    "position": "COO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 524,
                                    "sort": 0
                                },
                                {
                                    "name": "Jarett Lee",
                                    "position": "Marketing Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 525,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "P3-Trident",
                "studio_id": "96c03483-b8c6-4fd9-a25b-2215e4b7a9f1",
                "platform": [
                    "mobile",
                    "pc",
                    "console"
                ],
                "studio": "Sharkmob",
                "project_id": "19",
                "priority": "Key",
                "genre": [
                    "RPG/Action"
                ],
                "expanded_genre": "Action MMORPG",
                "business_model": [
                    "P+MTX"
                ],
                "image_name": "",
                "developer": "Sharkmob",
                "publisher": "Unknown",
                "ip": "New",
                "tagline": "An action RPG with online and cooperative elements, setting on everchanging isles. Fight bosses and conquer enemies to grow your reputation among the lands of Atlantis.",
                "ambition": "",
                "similar_games": [
                    {
                        "id": "e6411fefc5658304a2539793433ae4337",
                        "unified_id": "e6411fefc5658304a2539793433ae4337",
                        "name": "Path of Exile",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/c89fa0a65d4a4cfa7446251528b45c4f.jpg",
                        "game_type": 3,
                        "genre": "Adventure|Role Playing|Role-Playing|RPG|Action|Role-playing (RPG)|Hack and slash/Beat 'em up|Action RPG",
                        "sub_genre": "",
                        "country_en_name": "New Zealand",
                        "country_en_name_abbr": "NZ",
                        "region": "Oceania",
                        "publisher_name": "Grinding Gear Games",
                        "developer_name": "Grinding Gear Games"
                    }
                ],
                "current_headcounts": "15+",
                "launch_date": "20261231",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20220601",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "19",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20221206",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "19",
                        "child": null
                    },
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20240131",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "19",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20241130",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "19",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20260331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "19",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20260731",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "19",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20261231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "19",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "748",
                                    "id": 414,
                                    "sort": 0
                                },
                                {
                                    "name": "tomoconnor",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1035",
                                    "id": 493,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Studio Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 494,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 332,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Wardogs",
                "studio_id": "4306c6a0-b2f6-4194-3896-fabd3350d103",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Splash Damage",
                "project_id": "75",
                "priority": "Top 10",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Splash Damage",
                "ip": "New",
                "tagline": "Shooter, Community Mode",
                "ambition": "LARGE OPEN WORLD MIL-SIM SANDBOX",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20260630",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20230627",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "75",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20250313",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "75",
                        "child": null
                    },
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20251031",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "75",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20260630",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "75",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "748",
                                    "id": 512,
                                    "sort": 0
                                },
                                {
                                    "name": "ivandavies",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1033",
                                    "id": 513,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 514,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "D5-Scramble",
                "studio_id": "1a195f56-ffe9-c556-14f1-4f731b44f300",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Sumo Digital",
                "project_id": "74",
                "priority": "Top 10",
                "genre": [
                    "Casual"
                ],
                "expanded_genre": "",
                "business_model": [
                    "Premium",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "IEGG",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20230630",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "74",
                        "child": null
                    },
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20240131",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "74",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20240331",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "74",
                        "child": [
                            {
                                "milestone": "-EA Review",
                                "milestone_node": "",
                                "date": "20241025",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "74"
                            }
                        ]
                    },
                    {
                        "milestone": "Early Access Launch",
                        "milestone_node": "",
                        "date": "20241231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "74",
                        "child": null
                    },
                    {
                        "milestone": "PC Full Launch",
                        "milestone_node": "",
                        "date": "20250131",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "74",
                        "child": null
                    },
                    {
                        "milestone": "Console/Mobile Launch",
                        "milestone_node": "",
                        "date": "20250625",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "74",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "748",
                                    "id": 515,
                                    "sort": 0
                                },
                                {
                                    "name": "tomoconnor",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1035",
                                    "id": 516,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 517,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "S3-Ulysses",
                "studio_id": "4306c6a0-b2f6-4194-3896-fabd3350d103",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Splash Damage",
                "project_id": "21",
                "priority": "Middle",
                "genre": [
                    "SOC"
                ],
                "expanded_genre": "SOC shooter",
                "business_model": [
                    "P+MTX"
                ],
                "image_name": "",
                "developer": "Splash Damage",
                "publisher": "Splash Damage",
                "ip": "New",
                "tagline": "The perils of space, hostile player colonists, and an unpredictable rogue AI running the spacecraft all conspire to kill you as they battle to survive on board the ULYSSES",
                "ambition": "Bring AAA & sci-fi to the AA survival genre\nTarget #1 SOC that brings in new audience with shooter mechanics",
                "similar_games": [
                    {
                        "id": "efac98152ddaf5d5d08d6286e2626e83c",
                        "unified_id": "efac98152ddaf5d5d08d6286e2626e83c",
                        "name": "Fortnite",
                        "cover": "https://ogdb-cdn.intlgame.cn/pic_by_handle/fortnite pc.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Battle Royale|Tactical Third Person Shooter|Role-playing (RPG)|Action|Third-Person Shooter|Survival|Strategy",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Epic Games",
                        "developer_name": "Tencent Games"
                    }
                ],
                "current_headcounts": "80",
                "launch_date": "20260531",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20220214",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "21",
                        "child": null
                    },
                    {
                        "milestone": "First Playable gate materials ready",
                        "milestone_node": "",
                        "date": "20220930",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "21",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20241231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "21",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20250430",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "21",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20250930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "21",
                        "child": null
                    },
                    {
                        "milestone": "Early Access Launch",
                        "milestone_node": "",
                        "date": "20251130",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "21",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20260531",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "21",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 382,
                                    "sort": 0
                                },
                                {
                                    "name": "ivandavies",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1033",
                                    "id": 452,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 492,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 299,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Anth Rodgers",
                                    "position": "Executive Producer",
                                    "link": "",
                                    "user_id": "",
                                    "id": 138,
                                    "sort": 0
                                },
                                {
                                    "name": "Dr. Kay Ross",
                                    "position": "Technical Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 139,
                                    "sort": 0
                                },
                                {
                                    "name": "Lance Winter",
                                    "position": "Creative Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 140,
                                    "sort": 0
                                },
                                {
                                    "name": "Nikolay Georgiev",
                                    "position": "Art Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 141,
                                    "sort": 0
                                },
                                {
                                    "name": "Tim Rose",
                                    "position": "Production Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 142,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "YES",
                "studio_id": "c69f9642-bda6-4347-92e7-e0d33d27fd00",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Yager",
                "project_id": "83",
                "priority": "Low",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "Extraction Shooter",
                "business_model": [
                    "Premium",
                    "Microtransaction"
                ],
                "image_name": "",
                "developer": "",
                "publisher": "Yager",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20231218",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "83",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20240307",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "83",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": null
        },
        {
            "detail": {
                "project_name": "T2 - TAM",
                "studio_id": "cc69ad75-2fa2-a3f5-5be2-40fa618f141b",
                "platform": [
                    "mobile",
                    "pc",
                    "console"
                ],
                "studio": "Tequila Works",
                "project_id": "58",
                "priority": "Middle",
                "genre": [
                    "RPG",
                    "Action"
                ],
                "expanded_genre": "Adventure, narrative-driven, puzzle",
                "business_model": [
                    "Premium",
                    "Downloadle Content"
                ],
                "image_name": "",
                "developer": "Tequila Works",
                "publisher": "Unknown",
                "ip": "Existing",
                "tagline": "Open world single-player action RPG, narrative-driven adventure",
                "ambition": "Develop an ambitious action-adventure single player narrative driven game building on Tequila Works know-how and history.",
                "similar_games": [
                    {
                        "id": "eb0b533904195b362a4c991b137283082",
                        "unified_id": "eb0b533904195b362a4c991b137283082",
                        "name": "Shadow of the Colossus",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/f14fba1ece51c79a7324d965270d5f8b.jpg",
                        "game_type": 3,
                        "genre": "Action-Adventure|Action Adventure|Adventure|Platform|Puzzle",
                        "sub_genre": "",
                        "country_en_name": "Japan",
                        "country_en_name_abbr": "JP",
                        "region": "Japan",
                        "publisher_name": "Sony Interactive Entertainment",
                        "developer_name": "Bluepoint Games"
                    },
                    {
                        "id": "ebb419e1b4d75f5139c40a7c32eb67a94",
                        "unified_id": "ebb419e1b4d75f5139c40a7c32eb67a94",
                        "name": "Ōkami",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/a243ecf5d5deb991436a2c78708c8568.jpg",
                        "game_type": 2,
                        "genre": "Adventure",
                        "sub_genre": "",
                        "country_en_name": "Japan",
                        "country_en_name_abbr": "JP",
                        "region": "Japan",
                        "publisher_name": "Capcom",
                        "developer_name": "Tose|Clover Studio|Ready at Dawn"
                    }
                ],
                "current_headcounts": "1-5",
                "launch_date": "20251231",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch materials delivery",
                        "milestone_node": "",
                        "date": "20221020",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "58",
                        "child": null
                    },
                    {
                        "milestone": "Pitch review meeting",
                        "milestone_node": "",
                        "date": "20221020",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "58",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20230505",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "58",
                        "child": null
                    },
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20240331",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "58",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20240930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "58",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20241231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "58",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20251231",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "58",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 367,
                                    "sort": 0
                                },
                                {
                                    "name": "Jon Lawrence",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "",
                                    "id": 437,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "P2-Rebirth EXOBORNE",
                "studio_id": "96c03483-b8c6-4fd9-a25b-2215e4b7a9f1",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Sharkmob",
                "project_id": "18",
                "priority": "Top 5",
                "genre": [
                    "SOC",
                    "Shooter",
                    "MMO"
                ],
                "expanded_genre": "multiplayer survival third-person shooter",
                "business_model": [
                    "P+MTX"
                ],
                "image_name": "",
                "developer": "Sharkmob",
                "publisher": "IEGG",
                "ip": "New",
                "tagline": "\"Experience our world ruined by climate disaster. Join your friends online as you craft, build, and fight to survive.\"",
                "ambition": "First AAA SOC + MMO in a world ruined by climate disaster",
                "similar_games": [
                    {
                        "id": "e6667dd09505ddd44cf4ef980782e632c",
                        "unified_id": "e6667dd09505ddd44cf4ef980782e632c",
                        "name": "Warframe",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/424bf8ea992c376b5524ef24bf31c9bf.jpg",
                        "game_type": 3,
                        "genre": "Action|Misc|Role-playing (RPG)|Adventure|Role Playing|Shooter|Third-Person Shooter|Third Person Shooter",
                        "sub_genre": "",
                        "country_en_name": "Canada",
                        "country_en_name_abbr": "CA",
                        "region": "North America",
                        "publisher_name": "Digital Extremes",
                        "developer_name": "Digital Extremes"
                    },
                    {
                        "id": "ed09fe5188fb2059d33eb6ba04958bbb8",
                        "unified_id": "ed09fe5188fb2059d33eb6ba04958bbb8",
                        "name": "Tom Clancy's The Division",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/adf676d00f4dc27d82de8da0da881c63.jpg",
                        "game_type": 2,
                        "genre": "Shooter|Adventure|Action Games|Tactical Third Person Shooter|Action|Third-Person Shooter|Survival",
                        "sub_genre": "",
                        "country_en_name": "France",
                        "country_en_name_abbr": "FR",
                        "region": "Europe",
                        "publisher_name": "Ubisoft",
                        "developer_name": "Massive Entertainment"
                    }
                ],
                "current_headcounts": "292",
                "launch_date": "20251130",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20201203",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20210804",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": null
                    },
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20211110",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20230112",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20240112",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20240419",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": [
                            {
                                "milestone": "Public Test 1 - Retention",
                                "milestone_node": "",
                                "date": "20240530",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "18"
                            },
                            {
                                "milestone": "Public Test 2 - Retention",
                                "milestone_node": "",
                                "date": "20240730",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "18"
                            }
                        ]
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20240923",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": [
                            {
                                "milestone": "Chapter 2 Release",
                                "milestone_node": "",
                                "date": "20241230",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "18"
                            }
                        ]
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20251130",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "18",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "748",
                                    "id": 408,
                                    "sort": 0
                                },
                                {
                                    "name": "tomoconnor",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1035",
                                    "id": 479,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 490,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "Ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 326,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                },
                {
                    "organization": "studio",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "Fredrik Rundqvist",
                                    "position": "CEO",
                                    "link": "",
                                    "user_id": "",
                                    "id": 101,
                                    "sort": 0
                                },
                                {
                                    "name": "Johan Pfannenstill",
                                    "position": "Acting Production Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 102,
                                    "sort": 0
                                },
                                {
                                    "name": "Petter Mannerfelt\n",
                                    "position": "Creative Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 103,
                                    "sort": 0
                                },
                                {
                                    "name": "Rodrigo Cortes",
                                    "position": "Art Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 104,
                                    "sort": 0
                                },
                                {
                                    "name": "Martin Hultberg",
                                    "position": "IP Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 105,
                                    "sort": 0
                                },
                                {
                                    "name": "Niklas Westberg",
                                    "position": "Technical Director",
                                    "link": "",
                                    "user_id": "",
                                    "id": 106,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Path of Exile Mobile",
                "studio_id": "c522b63b-cd72-4cc1-8454-71337edd0db1",
                "platform": [
                    "mobile"
                ],
                "studio": "Grinding Gear Games",
                "project_id": "12",
                "priority": "Middle",
                "genre": [
                    "RPG",
                    "Action"
                ],
                "expanded_genre": "ARPG",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Project finance",
                "publisher": "Proxima Beta",
                "ip": "Existing",
                "tagline": "An experimental version of Path of Exile for mobile devices, a true mobile successor of Diablo's core gameplay",
                "ambition": "Expand the IP to a larger user group in mobile than existing PC+Console scale",
                "similar_games": [
                    {
                        "id": "u5eeaaef6e4cd86c631343b7ec3d26a22",
                        "unified_id": "u5eeaaef6e4cd86c631343b7ec3d26a22",
                        "name": "Torchlight: Infinite",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/274e88652f3d116ac912019068ff601d.jpg",
                        "game_type": 1,
                        "genre": "RPG",
                        "sub_genre": "Action RPG",
                        "country_en_name": "China",
                        "country_en_name_abbr": "CN",
                        "region": "China(Mainland)",
                        "publisher_name": "X.D. Network",
                        "developer_name": "X.D. Network"
                    },
                    {
                        "id": "u4b7afd39407c450c34cc643ed04bb114",
                        "unified_id": "u4b7afd39407c450c34cc643ed04bb114",
                        "name": "Diablo Immortal",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/655d1034c981e5e97af994b4ecedf80b.jpg",
                        "game_type": 1,
                        "genre": "RPG",
                        "sub_genre": "MMORPG",
                        "country_en_name": "",
                        "country_en_name_abbr": "",
                        "region": "",
                        "publisher_name": "NetEase|Blizzard Entertainment",
                        "developer_name": "NetEase|网易联合暴雪开发|Blizzard Entertainment"
                    }
                ],
                "current_headcounts": "10",
                "launch_date": "20240630",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20210930",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "12",
                        "child": null
                    },
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20221202",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "12",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20230509",
                        "date_confirm": "1",
                        "gate_review": "Cancel",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "12",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20230929",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "12",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta",
                        "milestone_node": "",
                        "date": "20231229",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "12",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20240630",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "12",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "bubuqian",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 386,
                                    "sort": 0
                                },
                                {
                                    "name": "reginagao",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "316",
                                    "id": 456,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "celinaluo",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "583",
                                    "id": 303,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Sniper Island",
                "studio_id": "b9eb2822-306d-491d-99a8-45dbf24e4e27",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Rebellion Developments Ltd",
                "project_id": "69",
                "priority": "Middle",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Rebellion",
                "publisher": "Tencent",
                "ip": "New",
                "tagline": "",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "",
                "launch_date": "20241104",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Document",
                        "milestone_node": "",
                        "date": "20220921",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice 1",
                        "milestone_node": "",
                        "date": "20230222",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice 2",
                        "milestone_node": "",
                        "date": "20230516",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice 3",
                        "milestone_node": "",
                        "date": "20231026",
                        "date_confirm": "1",
                        "gate_review": "Redo",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": [
                            {
                                "milestone": "-Vertical Slice 3 Redo",
                                "milestone_node": "",
                                "date": "20240216",
                                "date_confirm": "0",
                                "gate_review": "",
                                "key_feedback": "",
                                "action_items": null,
                                "approval_budget": "",
                                "milestone_date_reason": "",
                                "materials": null,
                                "milestone_node_start_date": "",
                                "project_id": "69"
                            }
                        ]
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20240311",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta",
                        "milestone_node": "",
                        "date": "20240520",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": null
                    },
                    {
                        "milestone": "Early Access",
                        "milestone_node": "",
                        "date": "20240715",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20241104",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "69",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "gramxu",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 381,
                                    "sort": 0
                                },
                                {
                                    "name": "Haleymzhang",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "12058",
                                    "id": 451,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "-",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "",
                                    "id": 298,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "Darktide (Warhammer 40,000)",
                "studio_id": "80b39a2b-6b55-4066-868b-029bbd2bc8b1",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Fatshark AB",
                "project_id": "6",
                "priority": "Middle",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "Co-op action/shooter",
                "business_model": [
                    "P+DLC"
                ],
                "image_name": "",
                "developer": "Self developed",
                "publisher": "Fatshark AB",
                "ip": "Licensed",
                "tagline": "",
                "ambition": "Create a Metacritic score of 85 for the spiritual sequel of Vermintide 2 under Warhammer 40K IP High-quality restoration of 40K, the Top 1 product in the Warhammer IP series",
                "similar_games": [
                    {
                        "id": "ee5fe56d8d433819bc746c96bc7fa3984",
                        "unified_id": "ee5fe56d8d433819bc746c96bc7fa3984",
                        "name": "Back 4 Blood",
                        "cover": "https://ogdb-cdn.intlgame.cn/image/c76596d020926d7330beb0f5bef65be4.jpg",
                        "game_type": 3,
                        "genre": "First-Person Shooter|Action|FPS|Shooter|Adventure",
                        "sub_genre": "",
                        "country_en_name": "United States",
                        "country_en_name_abbr": "US",
                        "region": "North America",
                        "publisher_name": "Warner Bros. Interactive Entertainment",
                        "developer_name": "Turtle Rock Studios"
                    }
                ],
                "current_headcounts": "173",
                "launch_date": "20221130",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20210420",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "6",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20211022",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "6",
                        "child": null
                    },
                    {
                        "milestone": "Closed Beta",
                        "milestone_node": "",
                        "date": "20221014",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "6",
                        "child": null
                    },
                    {
                        "milestone": "Open Beta",
                        "milestone_node": "",
                        "date": "20221117",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "6",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20221130",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "6",
                        "child": null
                    },
                    {
                        "milestone": "Post Launch",
                        "milestone_node": "",
                        "date": "20230930",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "6",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "gramxu",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 390,
                                    "sort": 0
                                },
                                {
                                    "name": "victorshen",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "208",
                                    "id": 460,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "gracehhuang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "11050",
                                    "id": 307,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "D1-Infinity",
                "studio_id": "1a195f56-ffe9-c556-14f1-4f731b44f300",
                "platform": [
                    "mobile",
                    "pc",
                    "console"
                ],
                "studio": "Sumo Digital",
                "project_id": "57",
                "priority": "Key",
                "genre": [
                    "RPG",
                    "Action"
                ],
                "expanded_genre": "Openworld, adventure",
                "business_model": [
                    "F2P"
                ],
                "image_name": "",
                "developer": "Sumo Digital",
                "publisher": "Unknown",
                "ip": "Existing",
                "tagline": "A unique AAA openworld adventure for all ages. An action RPG set in a diverse, beautiful and connected world.",
                "ambition": "",
                "similar_games": null,
                "current_headcounts": "140",
                "launch_date": "",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Pitch",
                        "milestone_node": "",
                        "date": "20220627",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "57",
                        "child": null
                    },
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20230323",
                        "date_confirm": "1",
                        "gate_review": "Pass",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "57",
                        "child": null
                    },
                    {
                        "milestone": "Concept/Prototype",
                        "milestone_node": "",
                        "date": "20240330",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "57",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 383,
                                    "sort": 0
                                },
                                {
                                    "name": "Tom O'Connor",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "",
                                    "id": 453,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 300,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        },
        {
            "detail": {
                "project_name": "S1-Fortress (Transformers)",
                "studio_id": "4306c6a0-b2f6-4194-3896-fabd3350d103",
                "platform": [
                    "pc",
                    "console"
                ],
                "studio": "Splash Damage",
                "project_id": "20",
                "priority": "Key",
                "genre": [
                    "Shooter"
                ],
                "expanded_genre": "Co-op action/shooter",
                "business_model": [
                    "Premium"
                ],
                "image_name": "",
                "developer": "Splash Damage",
                "publisher": "Splash Damage",
                "ip": "Licensed",
                "tagline": "AAA open-world co-op looter shooter",
                "ambition": "Splash Damage's strategic project to strengthen GaaS, F2P & publishing capabilities, and build gen 2.0 backend technology",
                "similar_games": null,
                "current_headcounts": "150",
                "launch_date": "20250930",
                "cover": "",
                "milestones": [
                    {
                        "milestone": "Concept/KOM",
                        "milestone_node": "",
                        "date": "20210701",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "20",
                        "child": null
                    },
                    {
                        "milestone": "First Playable",
                        "milestone_node": "",
                        "date": "20211213",
                        "date_confirm": "1",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "20",
                        "child": null
                    },
                    {
                        "milestone": "Vertical Slice",
                        "milestone_node": "",
                        "date": "20240430",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "20",
                        "child": null
                    },
                    {
                        "milestone": "Alpha",
                        "milestone_node": "",
                        "date": "20250131",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "20",
                        "child": null
                    },
                    {
                        "milestone": "Beta",
                        "milestone_node": "",
                        "date": "20250430",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "20",
                        "child": null
                    },
                    {
                        "milestone": "Official Launch",
                        "milestone_node": "",
                        "date": "20250930",
                        "date_confirm": "2",
                        "gate_review": "",
                        "key_feedback": "",
                        "action_items": null,
                        "approval_budget": "",
                        "milestone_date_reason": "",
                        "materials": null,
                        "milestone_node_start_date": "",
                        "project_id": "20",
                        "child": null
                    }
                ]
            },
            "update": null,
            "stakeholder": [
                {
                    "organization": "iegg",
                    "owner_list": [
                        {
                            "project_owner": "development",
                            "list": [
                                {
                                    "name": "petesmith",
                                    "position": "business_owner",
                                    "link": "",
                                    "user_id": "",
                                    "id": 407,
                                    "sort": 0
                                },
                                {
                                    "name": "ivandavies",
                                    "position": "pteam_contact",
                                    "link": "",
                                    "user_id": "1033",
                                    "id": 478,
                                    "sort": 0
                                },
                                {
                                    "name": "jonlawrence",
                                    "position": "Development Director",
                                    "link": "",
                                    "user_id": "1034",
                                    "id": 497,
                                    "sort": 0
                                }
                            ]
                        },
                        {
                            "project_owner": "finance",
                            "list": [
                                {
                                    "name": "ericxzhang",
                                    "position": "finance_contact",
                                    "link": "",
                                    "user_id": "1380",
                                    "id": 325,
                                    "sort": 0
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    ],
    "req": null
}
	`
}
