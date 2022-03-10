package test_data

const JsGetAdvMessagesI = `
{
  "GetAdvMessages": {
    "Advertisers": {
      "ID": "1"
    },
    "AdvertisingMessageIDs": [
      {
        "ID": "1"
      },
      {
        "ID": "2"
      }
    ],
    "Aspects": {
      "ID": "2"
    },
    "CreationDateEnd": "2019-03-02",
    "CreationDateStart": "2019-03-01",
    "FillMaterialTags": "true",
    "attributes": {
      "xsi": "http://www.w3.org/2001/XMLSchema-instance"
    }
  }
}`

const JsGetAdvMessagesO = `
{
 "GetAdvMessagesResult": {
  "AdvertisingMessagesData": [
   {
    "Row": {
     "AdtID": "700061957",
     "AdtName": "Название рекламодателя",
     "AspectID": "203",
     "AspectName": "4x3",
     "BrandID": "44362",
     "BrandName": "Название ТБ",
     "DoubleAdvertiser": "false",
     "FilmDur": "20",
     "FilmID": "861555",
     "FilmName": "Название ролика",
     "FilmVersion": "",
     "GroupID": "888064",
     "HasBroadcastMaterials": "true",
     "HasPreviewMaterials": "true",
     "HasSpots": "false",
     "ProdClassID": "703",
     "ffoaAllocated": "1",
     "ffoaLawAcc": "0"
    }
   },
   {
    "Row": {
     "AdtID": "700072339",
     "AdtName": "Название рекламодателя",
     "AspectID": "203",
     "AspectName": "4x3",
     "BrandID": "68436",
     "BrandName": "Название ТБ",
     "DoubleAdvertiser": "false",
     "FilmDur": "10",
     "FilmID": "861557",
     "FilmName": "Название ролика",
     "FilmVersion": "Версия ролика",
     "GroupID": "888064",
     "HasBroadcastMaterials": "true",
     "HasPreviewMaterials": "true",
     "HasSpots": "false",
     "ProdClassID": "426",
     "ffoaAllocated": "1",
     "ffoaLawAcc": "0"
    }
   }
  ],
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsGetBudgetsI = `
{
 "GetBudgets": {
  "AdvertiserList": {
   "ID": "700064621"
  },
  "ChannelList": {
   "ID": "1018574"
  },
  "EndMonth": "20180115",
  "SellingDirectionID": "21",
  "StartMonth": "201801"
 }
}`

const JsGetBudgetsO = `
{
 "responseBudget": {
  "BudgetList": {
   "m": {
    "AdtID": "17845",
    "AdtName": "Название рекламодателя",
    "AgrID": "17983",
    "AgrName": "Название договора",
    "Budget": "0.0",
    "CmpName": "Название компании",
    "CnlID": "1018574",
    "CnlName": "СТС",
    "CoordCost": "0.0",
    "Cost": "0.0",
    "DealChannelStatus": "1",
    "FixPercent": "0.0",
    "FixPercentPrime": "0.0",
    "FloatPercent": "0.0",
    "FloatPercentPrime": "0.0",
    "GRP": "0.0",
    "GRPFix": "0.0",
    "GRPWithoutKF": "0.0",
    "InventoryUnitDuration": "30",
    "Month": "201801",
    "Quality": {
     "Item": {
      "Percent": "0.55",
      "PercentPrime": "0.6",
      "RankID": "1"
     }
    },
    "TP": "Название точки родаж"
   }
  },
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsGetChannelsI = `
{
 "GetChannels": {
  "SellingDirectionID": "21"
 }
}`

const JsGetChannelsO = `
{
 "responseGetChannels": [
  {
   "Channel": {
    "ID": "1",
    "IsDisabled": "1",
    "MainChnl": "1",
    "ShortName": "Канал Дисней",
    "bcpCentralID": "81",
    "bcpName": "Москва",
    "cnlCentralID": "1"
   }
  },
  {
   "Channel": {
    "ID": "2",
    "IsDisabled": "0",
    "MainChnl": "2",
    "ShortName": "Канал Пятница",
    "bcpCentralID": "81",
    "bcpName": "Москва",
    "cnlCentralID": "0"
   }
  }
 ]
}`

const JsGetCustomersWithAdvertisersI = `
{
 "GetCustomersWithAdvertisers": {
  "SellingDirectionID": "23"
 }
}
`

const JsGetCustomersWithAdvertisersO = `
{
 "responseGetCustomersWithAdvertisers": {
  "Advertisers": [
   {
    "item": {
     "ID": "1",
     "Name": "SOME ADVERTISER"
    }
   },
   {
    "item": {
     "ID": "2",
     "Name": "ANOTHER ADVERTISER"
    }
   }
  ],
  "Data": [
   {
    "item": {
     "AdvID": "1",
     "CustID": "21"
    }
   },
   {
    "item": {
     "AdvID": "2",
     "CustID": "21"
    }
   }
  ]
 }
}`

const JsGetDeletedSpotInfoI = `
{
 "GetDeletedSpotInfo": {
  "Agreements": {
   "ID": "1"
  },
  "DateEnd": "2019-12-11T12:00:00",
  "DateStart": "2019-12-11T00:00:00"
 }
}`

const JsGetDeletedSpotInfoO = `
{
 "responseGetDeletedSpotInfo": {
  "Items": {
   "i": {
    "AffiliationType": "0",
    "AgrID": "0",
    "AgrName": "Название сделки",
    "BlockDate": "20191103",
    "BlockID": "180487969",
    "BlockNumber": "1",
    "BlockTime": "32837",
    "CnlName": "Первый",
    "CurrentAuctionBidValue": "0",
    "DeleteDateTime": "2019-12-11T09:29:05",
    "FilmDur": "45",
    "FilmID": "912037",
    "FilmName": "Название ролика",
    "FilmVersion": "Версия ролика",
    "OrdID": "484148",
    "OrdName": "Название заказа",
    "Position": "-3",
    "PrgName": "Бокс. Бой за титул чемпиона мира",
    "Reason": "1",
    "SpotID": "755086586"
   }
  },
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsGetRanksI = `
{
 "GetRanks": ""
}`

const JsGetRanksO = `
{
 "responseGetRanks": {
  "Ranks": {
   "Rank": {
    "Details": [
     {
      "Detail": {
       "EndDate": "2079-06-06T23:59:00",
       "OrderNo": "0",
       "SellingDirectionID": "18",
       "StartDate": "1900-01-01T00:00:00",
       "UsesAuction": "false"
      }
     },
     {
      "Detail": {
       "EndDate": "2020-01-01T00:00:00",
       "OrderNo": "0",
       "SellingDirectionID": "21",
       "StartDate": "1900-01-01T00:00:00",
       "UsesAuction": "true"
      }
     },
     {
      "Detail": {
       "EndDate": "2079-06-06T23:59:00",
       "OrderNo": "1",
       "SellingDirectionID": "21",
       "StartDate": "2020-01-01T00:00:00",
       "UsesAuction": "false"
      }
     },
     {
      "Detail": {
       "EndDate": "2079-06-06T23:59:00",
       "OrderNo": "0",
       "SellingDirectionID": "23",
       "StartDate": "1900-01-01T00:00:00",
       "UsesAuction": "false"
      }
     }
    ],
    "ID": "2",
    "Name": "Фикс"
   }
  },
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsGetSpotsI = `
{
 "GetSpots": {
  "AdtList": {
   "AdtID": "1"
  },
  "ChannelList": {
   "Cnl": "2",
   "Main": "1"
  },
  "EndDate": "20161016",
  "InclOrdBlocks": "1",
  "SellingDirectionID": "21",
  "StartDate": "20161010"
 }
}`

const JsGetSpotsO = `
{
 "ResponseSpot": {
  "OrdBlocks": {
   "obl": {
    "BlockID": "0",
    "IRate": "0",
    "OrdID": "0",
    "Rate": "0.01"
   }
  },
  "SpotList": {
   "s": {
    "AgrID": "0",
    "AllocationType": "0",
    "AtpID": "0",
    "BaseRating": "0.768",
    "BlockID": "0",
    "CommInMplID": "0",
    "CurrentAuctionBidValue": "",
    "DtpID": "",
    "FixPriority": "",
    "FloatPriority": "0",
    "IBaseRating": "0",
    "IRating": "0",
    "IsHumanBeing": "true",
    "MplID": "0",
    "OTS": "230.745",
    "OrdID": "0",
    "Positioning": "",
    "RankID": "3",
    "Rating": "0.768",
    "SpotBroadcastTime": "86063",
    "SpotFactBroadcastTime": "86063",
    "SpotID": "0",
    "SpotOrderNo": "1",
    "SpotReserve": "1",
    "SptDateL": "1032929",
    "TNSBlockID": "",
    "TNSSpotsID": "",
    "TgrID": "",
    "sptChnlPTR": "0"
   }
  },
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsAddSpotI = `
{
 "AddSpot": {
  "AuctionBidValue": "",
  "BlockID": "1",
  "FilmID": "1238114",
  "FixedPosition": "true",
  "Position": "",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsAddSpotO = `
{
 "responseAddSpot": {
  "b": {
   "AddSpotStart": "Jan  1 1900 12:00AM",
   "ErrorMessage": "",
   "Position": "",
   "Result": "1",
   "newSptID": "1"
  }
 }
}`

const JsDeleteSpotI = `
{
 "DeleteSpot": {
  "SpotID": "123"
 }
}`

const JsDeleteSpotO = `
{
 "responseDeleteSpot": {
  "b": {
   "ErrorMessage": "",
   "Result": "1"
  }
 }
}`

const JsChangeSpotI = `
{
 "ChangeSpots": {
  "FirstSpotID": "1",
  "SecondSpotID": "2"
 }
}`

const JsChangeSpotO = `
{
 "responseChangeSpots": {
  "b": {
   "ErrorMessage": "",
   "Result": "1"
  }
 }
}`

const JsSetSpotPositionI = `
{
 "SetSpotPosition": {
  "Distance": "2",
  "SpotID": "637435949"
 }
}`

const JsSetSpotPositionO = `{
 "SetSpotPositionResult": {
  "IsSuccess": "true",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsChangeFilmsI = `
{
 "ChangeFilms": [
  {
   "FakeSpotIDs": "1"
  },
  {
   "FakeSpotIDs": "2"
  },
  {
   "CommInMplIDs": "1"
  },
  {
   "CommInMplIDs": "2"
  }
 ]
}`

const JsChangeFilmsO = `
{
 "responseChangeFilm": {
  "b": {
   "ErrorMessage": "",
   "Result": "1",
   "SpotData": [
    {
     "Item": {
      "SplitMessageID": "2062725",
      "SpotID": "210448182"
     }
    },
    {
     "Item": {
      "SplitMessageID": "2062725",
      "SpotID": "210448182"
     }
    }
   ],
   "SpotIDs": [
    {
     "ID": "12345678"
    },
    {
     "ID": "12345679"
    }
   ]
  }
 }
}`

const JsAddMPlanI = `
{
 "AddMPlan": {
  "BrandID": "",
  "DateFrom": "2020-06-01",
  "DateTo": "2020-06-30",
  "MplCnlID": "1019882",
  "MplName": "Тест",
  "MultiSpotsInBlock": "false",
  "OrdID": "499499",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsAddMPlanO = `
{
 "responseAddMPlan": {
  "MplID": "16668506",
  "MplName": "Матч ТВ Ярославль Июнь 2020 фикс",
  "MplState": "2",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsAddMPlanFilmI = `
{
 "AddMPlanFilm": {
  "EndDate": "",
  "FilmID": "751900",
  "MplID": "14396424",
  "StartDate": "2018-12-24",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsAddMPlanFilmO = `
{
 "responseAddMPlanFilm": {
  "CommInMplID": "3231344",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsDeleteMPlanFilmI = `
{
 "DeleteMPlanFilm": {
  "CommInMplID": "9249649"
 }
}`

const JsDeleteMPlanFilmO = `
{
 "responseDeleteMPlanFilmResult": ""
}`

const JsChangeMPlanFilmPlannedInventoryI = `
{
 "ChangeMPlanFilmPlannedInventory": {
  "Data": [
   {
    "CommInMpl": {
     "ID": "3342386",
     "Inventory": "10"
    }
   },
   {
    "CommInMpl": {
     "ID": "3342386",
     "Inventory": "20"
    }
   },
   {
    "CommInMpl": {
     "ID": "93342386",
     "Inventory": "10"
    }
   }
  ]
 }
}`

const JsChangeMPlanFilmPlannedInventoryO = `
{
 "ChangeMPlanFilmPlannedInventoryResult": {
  "IsSucceded": "true",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsGetMPLansI = `
{
 "GetMPlans": {
  "AdtList": {
   "AdtID": "1"
  },
  "ChannelList": {
   "Cnl": "1018578"
  },
  "EndMonth": "201707",
  "IncludeEmpty": "true",
  "SellingDirectionID": "21",
  "StartMonth": "201707",
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsGetMPLansO = `
{
 "responseMPlans": {
  "MPlansList": {
   "m": {
    "AdtID": "700073097",
    "AdtName": "Название рекламодателя",
    "AdvID": "700073097",
    "AgrID": "12358",
    "AgrName": "Название сделки",
    "AllocationAllowed": "false",
    "AllocationType": "0",
    "AmountFact": "100.1",
    "AmountPlan": "100.1",
    "BrandID": "52091",
    "BrandName": "Название Товарного бренда",
    "CPPoffprime": "100",
    "CPPprime": "100",
    "CommInMplID": "1236308",
    "CommInMplPlanBudget": "7327.2931171718",
    "CommInMplPlanInventory": "6.6039551918569",
    "CommInMplPlanPrimeBudget": "7327.2931171718",
    "CommInMplPlanPrimeInventory": "6.6039551918569",
    "ContractBeg": "20170101",
    "ContractEnd": "20171231",
    "DateFrom": "20170701",
    "DateTo": "20170709",
    "DealChannelStatus": "1",
    "Discounts": [
     {
      "item": {
       "ApplicableToDeals": "false",
       "ApplyingTypeID": "10",
       "DicountEndDate": "",
       "DiscountFactor": "1.00000000000000",
       "DiscountStartDate": "",
       "DiscountTypeName": "За позиционирование 30",
       "IsManual": "false",
       "IsSpotPositionDependent": "true",
       "TypeID": "1",
       "ValueID": "1"
      }
     },
     {
      "item": {
       "AggregationMethodName": "",
       "ApplicableToDeals": "true",
       "ApplyingTypeID": "8",
       "DicountEndDate": "2017-07-31T00:00:00",
       "DiscountFactor": "0.80000000000000",
       "DiscountStartDate": "2017-07-01T00:00:00",
       "DiscountTypeName": "Сезонная ",
       "IsDiscountAggregate": "false",
       "IsManual": "false",
       "IsSpotPositionDependent": "false",
       "TypeID": "10",
       "ValueID": ""
      }
     }
    ],
    "DoubleAdvertiser": "true",
    "DtpID": "0",
    "DublSpot": "0",
    "FbrName": "Название Финансового бренда",
    "FilmDur": "20",
    "FilmDurKoef": "1.0",
    "FilmID": "655571",
    "FilmName": "Название Ролика",
    "FilmVersion": "Версия Ролика",
    "FixPriceAsFloat": "0",
    "FixPriority": "",
    "GRP": "2.2",
    "GRPShift": "0.5",
    "GroupID": "942772",
    "GrpFact": "2.2",
    "GrpPlan": "2.2",
    "GrpTotal": "2.2",
    "GrpTotalPrime": "13.2586340243",
    "HasReserve": "1",
    "Inventory": "0",
    "InventoryUnitDuration": "30",
    "InventoryUnits": "49",
    "MainOrderID": "",
    "MplCbrID": "123",
    "MplCbrName": "Название Товарного бренда из медиаплана",
    "MplCnlID": "1018578",
    "MplID": "13425963",
    "MplMonth": "20170701",
    "MplName": "Название медиаплана ",
    "MplState": "1",
    "Multiple": "",
    "OBDPos": "1",
    "OrdFrID": "848971",
    "OrdID": "80120",
    "OrdIsTriggered": [
     "0",
     "0"
    ],
    "OrdName": "Название Заказа",
    "PBACond": "Условия ПБА",
    "PBAObjID": "9335",
    "PrimeShare": "0.65",
    "ProdClassID": "615",
    "RankID": "3",
    "SplitMessageGroupID": "",
    "SumShift": "0.5",
    "TPName": "",
    "TgrID": "194",
    "TgrName": "все 18-44",
    "ffoaAllocated": "1",
    "ffoaLawAcc": "1",
    "ordBegDate": "20170501",
    "ordEndDate": "20170731",
    "ordManager": "ФИО"
   }
  },
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`

const JsGetProgramBreaksLight = `
{
 "responseProgramBreaksV2Light": {
  "BreakList": [
   {
    "b": {
     "Booked": {
      "attributes": {
       "RankID": "2",
       "VM": "175",
       "VR": "0"
      }
     },
     "attributes": {
      "BlockID": "117952183",
      "HasAucSpots": "false",
      "VM": "180",
      "VR": "0"
     }
    }
   },
   {
    "b": {
     "attributes": {
      "BlockID": "117952235",
      "HasAucSpots": "false",
      "VM": "180",
      "VR": "0"
     }
    }
   },
   {
    "b": {
     "Booked": [
      {
       "attributes": {
        "RankID": "2",
        "VM": "125",
        "VR": "0"
       }
      },
      {
       "attributes": {
        "RankID": "3",
        "VM": "10",
        "VR": "0"
       }
      }
     ],
     "attributes": {
      "BlockID": "117952251",
      "HasAucSpots": "false",
      "VM": "180",
      "VR": "0"
     }
    }
   }
  ],
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}
`

const JsGetProgramBreaksO = `
{
 "responseProgramBreaksV2": {
  "BlockForecast": {
   "B": {
    "BlockID": "1",
    "Fact": "",
    "Forecast": "0.2",
    "ForecastQuality": "good",
    "InternetForecast": "",
    "tgrID": "1"
   }
  },
  "BlockForecastTgr": {
   "Tgr": {
    "ID": "928",
    "Name": "Название ЦА"
   }
  },
  "BreakList": {
   "b": {
    "AucRate": "",
    "AuctionStepPrice": "",
    "AvailableAuctionVolume": "",
    "BlkAdvertTypePTR": "37",
    "BlkAuc": "0",
    "BlockDate": "2019-12-10",
    "BlockDistr": "0",
    "BlockNumber": "1",
    "BlockTime": "28620",
    "Booked": [
     {
      "attributes": {
       "RankID": "2",
       "VM": "155",
       "VR": "0"
      }
     },
     {
      "attributes": {
       "RankID": "3",
       "VM": "10",
       "VR": "0"
      }
     }
    ],
    "CnlID": "1021210",
    "CnlName": "Москва 24",
    "DLDate": "2019-12-05T13:00:00",
    "DLTrDate": "1900-01-01T00:00:00",
    "FactRateBase": "",
    "ForecastRateBase": "",
    "IsLocal": "1",
    "IsPrime": "0",
    "IsSpecialProject": "0",
    "IssDuration": "180",
    "IssID": "72950041",
    "IssTime": "28620",
    "NoRating": "1",
    "PrgBegMonthL": "28267",
    "PrgBegTimeL": "28320",
    "PrgEndMonthL": "28292",
    "PrgName": "Рекламный блок",
    "PrgNameShort": "Рекламный блок",
    "Pro2": "0",
    "ProID": "436837",
    "ProOriginalPTR": "",
    "ProgID": "674380",
    "RCID": "1021210",
    "RPID": "",
    "RateAll": "",
    "RootRCID": "1021210",
    "RootRPID": "674380",
    "SptOptions": "8",
    "TNSBlockFactDur": "",
    "TNSBlockFactID": "",
    "TNSBlockFactTime": "",
    "TgrID": "",
    "TgrName": "",
    "WeekDay": "2",
    "attributes": {
     "BlockID": "117952185",
     "VM": "180",
     "VR": "0"
    }
   }
  },
  "ProMaster": {
   "p": {
    "ProID": "215301",
    "PropName": "Тип",
    "PropValue": "Программа"
   }
  },
  "attributes": {
   "xsi": "http://www.w3.org/2001/XMLSchema-instance"
  }
 }
}`
