import json
import time
from itertools import repeat

import requests


if __name__ == "__main__":
    ip = "http://192.168.1.2:9000"
    # add groups
    group_infos = [{"groupName": "1DCS", "columnNames": "groupName,type,description,unit,source".split(",")}]
    requests.post(url=f"{ip}/group/addGroups", data=json.dumps(group_infos, ensure_ascii=False))
    # add items
    items = [{"itemName": "testItem1", "groupName": "1DCS", "type": "","description": "testItem1", "unit": "", "source": ""},
             {"itemName": "testItem2", "type": "","groupName": "1DCS", "description": "testItem2", "unit": "", "source": ""},
             {"itemName": "testItem3", "type": "","groupName": "1DCS", "description": "testItem3", "unit": "", "source": ""},
             {"itemName": "testItem4", "type": "","groupName": "1DCS", "description": "testItem4", "unit": "", "source": ""},
             {"itemName": "testItem5", "type": "","groupName": "1DCS", "description": "testItem5", "unit": "", "source": ""},
             {"itemName": "testItem6", "type": "","groupName": "1DCS", "description": "testItem6", "unit": "", "source": ""},
             {"itemName": "testItem7", "type": "","groupName": "1DCS", "description": "testItem7", "unit": "", "source": ""},
             {"itemName": "testItem8", "type": "","groupName": "1DCS", "description": "testItem8", "unit": "", "source": ""}]
    requests.post(f"{ip}/item/addItems", data=json.dumps({"groupName": "1DCS", "values": items}))
    # write realTime data without timeStamps
    item_values = [{"itemName": f"testItem{i}", "value": str(i * 100)} for i in range(1, 6)]
    requests.post(url=f"{ip}/data/batchWrite", data=json.dumps({"groupName": "1DCS", "itemValues": item_values, "withTimeStamp": False}))
    # write data with timestamp
    t = int(time.time())
    item_values = [{"itemName": f"testItem{i}", "value": str(i * 100), "timeStamp": str(t)} for i in range(6, 9)]
    requests.post(url=f"{ip}/data/batchWrite", data=json.dumps({"groupName": "1DCS", "itemValues": item_values, "withTimeStamp": True}))
    # get realTime data, return the latest updated data
    requests.post(url=f"{ip}/data/getRealTimeData", data=json.dumps({"itemNames": [f"testItem{i}" for i in range(1, 9)]}))
    # get historical data with timestamp
    requests.post(url=f"{ip}/data/getHistoricalDataWithStamp",
                  data=json.dumps({"itemNames": [f"testItem{i}" for i in range(6, 9)],
                                   "timeStamps": list(repeat([t], 3))}))
