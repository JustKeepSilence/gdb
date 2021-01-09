<template>
  <div>
    <el-row :gutter="10">
      <el-col :span="6" v-for="(item, index) in itemValues" :key="index" >
        <el-tooltip
          class="item"
          effect="dark"
          :content="item.tip"
          placement="top-start"
        >
          <el-card shadow="hover" style="cursor: pointer;height:120px">
            <el-row>
              <el-col :span="8" :xs='8' class="hidden-xs-only">
                <svg-icon :icon-class="item.name" />
              </el-col>
              <el-col :span="16" :xs='16'>
                <el-row> &nbsp; </el-row>
                <el-row>
                  <span class="card-font" v-if='!small'>{{ item.content }}</span>
                  <span class="card-font-mobile" v-else>{{ item.content }}</span>
                </el-row>
              </el-col>
            </el-row>
          </el-card>
        </el-tooltip>
      </el-col>
    </el-row>
    <br />
    <el-row>
      <el-card shadow="hover" style="height:500px">
        <div id="main" class="chart" v-if='!small'></div>
        <div id='main' class="chart_mobile" v-else></div>
      </el-card>
    </el-row>
    <el-row>
      <el-card shadow='hover' style="height:500px;margin-top: 20px">
         <div slot="header">
           <el-row>
              <span style="float:left;font-weight:bold"><i class="el-icon-link"></i>本项目基于但不限于以下开源项目</span>
           </el-row>
         </div>
         <div>
           <el-row v-for="(item, index) in sources" :key="index" >
              <el-link :href='item.url' style="float:left;font-size: 1.2em" >{{item.name}}: {{item.content}}</el-link><el-divider></el-divider>
           </el-row>
         </div>
      </el-card>
    </el-row>
  </div>
</template>
<script>
import {getDbInfo, getSpeedHistory} from '@/api'
import 'element-ui/lib/theme-chalk/display.css';
export default {
  name: "FirstPage",
  data: () => {
    return {
      itemValues: [
        { name: "ram", content: "123MB", tip: "内存使用信息" },
        { name: "writtenItems", content: "6321", tip: "写入Item个数" },
        { name: "currentTimeStamp", content: "123300000", tip: "系统时间戳" },
        { name: "speed", content: "12ms/6321", tip: "当前写入速率" },
      ],
      chartData: [],
      now: +new Date(1997, 9, 3),
      oneDay: 24 * 3600 * 1000,
      value: Math.random() * 1000,
      myChart: null,
      currentTime: +new Date(),
      xData:[], 
      rtSpeed: '',
      fontSize:'40px',
      small: false,
      sources: [{name: 'Go', url: 'https://github.com/golang/go', content: 'Google开源的支持并发的静态语言'}, {name: 'leveldb', url: 'https://github.com/google/leveldb', content: 'Google开源的基于C++的k-v数据库'},
                {name: 'goleveldb', url: 'https://github.com/syndtr/goleveldb', content: 'level的go语言版本,需要cgo'}, {name: 'gin', url: 'https://github.com/gin-gonic/gin', content: '基于go语言的并发网络框架'},
                {name: 'go-sqlite3', url: 'https://github.com/mattn/go-sqlite3', content: 'sqlite的go语言接口实现'}, {name: 'json-iterator', url: 'https://github.com/json-iterator/go', content:'代替内置json的高性能json解析器'},
                {name: 'goja', url: 'https://github.com/dop251/goja', content:'基于go的js容器'}, {name: 'cmap', url: 'https://github.com/orcaman/concurrent-map', content:'并发的map'},
                {name: 'vue', url: 'https://github.com/vuejs/vue', content:'渐进式的js框架'}, {name: 'element-ui', url: 'https://github.com/ElemeFE/element', content:'饿了么开源的基于vue的组件库'},
                {name: 'element-admin', url: 'https://github.com/PanJiaChen/vue-element-admin', content: '基于vue-element-ui的后端开源框架'}
      ]
    };
  },
  mounted() {
    if (document.body.clientWidth < 768){
       this.fontSize = '10px'
       this.small = true
    }
    document.querySelector(".el-main").style.backgroundColor = " #f0f2f5";
    this.render()
    this.myChart = this.$echarts.init(document.getElementById("main"));
    getSpeedHistory(JSON.stringify({
        startTimes: [this.currentTime - 10 * 60 * 1000],
        endTimes: [this.currentTime],
        intervals: [5]
    })).then(({data})=>{
        for(let i =0 ; i < data.speed[0].length;i++){
           this.xData.push(this.parseTime(new Date(parseInt(data.speed[0][i])* 1000)))
           this.chartData.push(parseFloat(data.speed[1][i].replace("ms", "")))
        }
        this.myChart.setOption({
      title: {
        text: "实时写入速率",
      },
      grid:{
      	 top:"50px",
         left:"50px",
         right:"15px",
         bottom:"50px"
	  },
      tooltip: {
        trigger: "axis",
        formatter: function (params) {
          return params[0].name + " " + params[0].value + "ms"      
        },
        axisPointer: {
          animation: false,
        },
      },
      xAxis: {
        data:this.xData,
        splitLine: {
          show: false,
        },
      },
      yAxis: {
        type: "value",
        boundaryGap: [0, "100%"],
        splitLine: {
          show: false,
        },
      },
      series: [
        {
          name: "模拟数据",
          type: "line",
          showSymbol: false,
          hoverAnimation: false,
          data: this.chartData,
          markLine: {
                silent: true,
                data: [{
                    yAxis: 30
                }]
            }
        },
      ],
    })
    })
    const _this = this;
    setInterval(function () {
      getDbInfo().then(({data})=>{
      for(let item of _this.itemValues){
        if (item.name== "ram"){
          item.content = data[item.name] + "MB"
        }else{
          item.content = data[item.name]
        }
        if(item.name=="speed"){
          _this.rtSpeed = parseInt(item.content.split("/")[0].replace("ms", ""))
        }
      }
      _this.xData.shift()
      _this.chartData.shift()
      _this.xData.push(_this.parseTime(new Date()))
      _this.chartData.push(_this.rtSpeed)
      _this.myChart.setOption({
        series: [
          {
            data: _this.chartData,
           
          },
        ],
        xAxis: {
        data: _this.xData,
        splitLine: {
            show: false
        }
        }
      })
    })  
    }, 1000)
  },
  methods: {
    randomData() {
      this.now = new Date(+this.now + this.oneDay);
      this.value = this.value + Math.random() * 21 - 10;
      return {
        name: this.now.toString(),
        value: [
          [
            this.now.getFullYear(),
            this.now.getMonth() + 1,
            this.now.getDate(),
          ].join("-"),
          Math.round(this.value),
        ],
      };
    },
    render() {
      getDbInfo().then(({data})=>{
      for(let item of this.itemValues){
        if (item.name== "ram"){
          item.content = data[item.name] + "MB"
        }else{
          item.content = data[item.name]
        }
        if(item.name=="speed"){
          this.rtSpeed = item.content
        }
      }
    })
    },
    parseTime(t) {
      return `${t.getFullYear()}-${
        t.getMonth() + 1 < 10 ? "0" + (t.getMonth() + 1) : t.getMonth() + 1
      }-${t.getDate() < 10 ? "0" + t.getDate() : t.getDate()} ${
        t.getHours() < 10 ? "0" + t.getHours() : t.getHours()
      }:${t.getMinutes() < 10 ? "0" + t.getMinutes() : t.getMinutes()}:${
        t.getSeconds() < 10 ? "0" + t.getSeconds() : t.getSeconds()
      }`;
    },
  },
};
</script>
<style scoped>

.el-icon-s-chartData {
  font-size: 60px;
}
.card-font {
  font-size: 40px;
  float: right;
  margin-top: 10px;
}
.card-font-mobile{
  font-size: 10px;
  margin-top: 10px;
  margin-left: -10px;
  line-height: 50px;
}
.chart{
  width: 100%; height: 400px
}
.chart_mobile{
  width: 300px; height: 400px
}
</style>