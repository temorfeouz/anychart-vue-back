<template>
  <div class="container" id="app">
    <h1 class="text-center">
    Atmosphere info
    </h1>
    <section class="chart-list">
      <!--<section class="chart-container">-->
        <!--<vue-anychart :options="lineOptions" ref="lineChart"></vue-anychart>-->
        <!--<button class="btn btn-primary" @click="add(getRandomData())">Add Series</button>-->
        <!--<button class="btn btn-danger" @click="remove" :disabled="lineSeriesCount == 0">Remove Series</button>-->
      <!--</section>-->

      <!--<section class="chart-container">-->
        <!--<vue-anychart :options="pieOptions" ref="pieChart"></vue-anychart>-->
        <!--<button class="btn btn-primary" @click.once="updatePieData" :disabled="pieDataIsModified == true">Update data-->
        <!--</button>-->
      <!--</section>-->

      <!--<section class="chart-container">-->
        <!--<vue-anychart :options="areaOptions" ref="areaChart"></vue-anychart>-->
        <!--<button class="btn btn-primary" @click.once="modifiedXAxis" :disabled="xAxisIsModified == true">Modify xAxis-->
        <!--</button>-->
      <!--</section>-->
      <input id="period" value="10m" v-model="period">
      <input id="startTime" value="" v-model="startTime">
      <input id="endTime" value="" v-model="endTime">
      <button class="btn btn-primary" @click="updateChartData">Update data
      </button>
      <section class="chart-container">
        <vue-anychart :options="CombineOptions" ref="combineChart"></vue-anychart>
      </section>
      <section class="chart-container">
        <vue-anychart :options="CombineOptions1" ref="combineChart1"></vue-anychart>
      </section>
    </section>

  </div>
</template>

<script>
/* eslint-disable */
/* use a function for the exact format desired... */
function ISODateString(d){
  function pad(n){return n<10 ? '0'+n : n}
 var ret=  d.getUTCFullYear()+'-'
    + pad(d.getUTCMonth()+1)+'-'
    + pad(d.getUTCDate())+'T'
   + pad(d.getUTCHours())+':'
     + pad(d.getUTCMinutes())+':'
     + pad(d.getUTCSeconds())+'Z'

  return ret
}

function timeConverter(UNIX_timestamp){
  var a = new Date(UNIX_timestamp*1000 );
  var months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
  var year = a.getFullYear();
  var month = months[a.getMonth()];
  var date = a.getDate();
  var hour = a.getHours();
  var min = a.getMinutes();
  var time = date + ' ' + month + ' ' + year + ' ' + hour + ':' + min  ;
  return time;
}
  import VueAnychart from './components/VueAnychart'
  import * as data from './data/data'
  import Anychart from 'anychart'
import { axios } from '@/plugins/axios'
  export default {
    name: 'App',
    components: {
      VueAnychart
    },
    data() {
      var ret = {
        period:"10m",
        startTime:"",
        endTime:"",
        Anychart: Anychart,
        // areaOptions: data.AreaData,
        // pieOptions: data.PieData,
        // lineOptions: data.LineData,
        CombineOptions: null,
        CombineOptions1: null,
        lineSeriesCount: 0,
        xAxisIsModified: false,
        pieDataIsModified: false
      }

      this.updateChartData()

      return ret
    },
    mounted() {
     var startdate = new Date()
      startdate.setDate(startdate.getDate()-10)
      this.$data.startTime=ISODateString(startdate)
      this.$data.endTime=ISODateString(new Date())


      this.lineSeriesCount = this.$refs.lineChart.chart.getSeriesCount()
    },
    methods: {
      remove() {
        this.$refs.lineChart.removeSeries()
        this.lineSeriesCount--
      },
      add(data) {
        this.$refs.lineChart.addSeries(data)
        this.lineSeriesCount++
      },
      removeAll() {
        this.$refs.lineChart.removeAllSeries()
        this.lineSeriesCount = 0
      },
      modifiedXAxis() {
        let xAxis = this.$refs.areaChart.chart.xAxis(0);
        xAxis.orientation('top');
        xAxis.labels().format(function () {
          return this.value.slice(0, 3)
        });

        this.xAxisIsModified = true;
      },



      updateChartData() {
       var newPeriod=""
        var startdate = new Date()
        startdate.setDate(startdate.getDate()-10)
        var startTime=ISODateString(startdate)
        var endTime=ISODateString(new Date())
        if (this.$data!=undefined){
         console.log("!!",this.$data.period)
         if (this.$data.period!=""){
           newPeriod='&step='+this.$data.period
         }
          startTime=this.$data.startTime
          endTime=this.$data.endTime
        }
        // newPeriod=""
        var url ='http://192.168.1.253:8428/api/v1/query_range?start='+startTime+'&end='+endTime+''+newPeriod+'&query={__name__=~"measurement.*"}'
       //'http://192.168.1.253:8428/api/v1/export?match={__name__=~%22measurement.*%22}'
        console.log(url)
        axios.get(url).then(
          response =>  {

            this.$data.CombineOptions= {
              'chart': {
                'title': 'Temp + humidity',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [ {
                  'seriesType': 'spline',
                  'name': 'Humidity',
                  'normal': {'stroke': {'color': 'blue', 'thickness': 2.5}},
                  'data': [
                  ]
                },
                  {
                    'seriesType': 'spline',
                    'name': 'Temperature',
                    'normal': {'stroke': {'color': 'red', 'thickness': 2.5}},
                    'data': [
                    ]
                  }
                ],
                'type': 'column'
              }
            }

            this.$data.CombineOptions1= {
              'chart': {
                'title': 'Gas + pressure',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [  {
                  'seriesType': 'spline',
                  'name': 'Gas',
                  'normal': {'stroke': {'color': 'black', 'thickness': 2.5}},
                  'data': [
                  ]
                },{
                  'seriesType': 'spline',
                  'name': 'Preassure',
                  'normal': {'stroke': {'color': 'green', 'thickness': 2.5}},
                  'data': [
                  ]
                }
                ],
                'type': 'column'
              }
            }

            response.data.data.result.forEach(v=>{
                if (v.metric.__name__=="measurement_humidity"){
                  for ( var i=0;i<v.values.length;i++){
                    if (v.values[i][1]==0){continue}
                    this.$data.CombineOptions.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': v.values[i][1]})
                  }
                }
                if (v.metric.__name__=="measurement_temperature"){
                  for ( var i=0;i<v.values.length;i++){
                    if (v.values[i][1]< -60){continue}
                    this.$data.CombineOptions.chart.series[1].data.push({'x': timeConverter(v.values[i][0]), 'value': v.values[i][1]})
                  }
                }
                if (v.metric.__name__=="measurement_gas"){
                  for ( var i=0;i<v.values.length;i++){
                    if (v.values[i][1]==0){continue}
                    this.$data.CombineOptions1.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': v.values[i][1]})
                  }
                }
                if (v.metric.__name__=="measurement_preassure"){
                  for ( var i=0;i<v.values.length;i++){
                    if (v.values[i][1]==0){continue}
                    this.$data.CombineOptions1.chart.series[1].data.push({'x': timeConverter(v.values[i][0]), 'value': v.values[i][1]})
                  }
                }
                          })


          }).catch((error) => alert("ERRRR_"+error));

        // this.pieDataIsModified = true;
      },
      getRandomData() {
        return [
          {'x': 'January', 'value': this.getRandomInt(1, 15) * 1000},
          {'x': 'February', 'value': this.getRandomInt(3, 18) * 1000},
          {'x': 'March', 'value': this.getRandomInt(2, 15) * 1000},
          {'x': 'April', 'value': this.getRandomInt(1, 18) * 1000},
          {'x': 'May', 'value': this.getRandomInt(3, 18) * 1000}
        ]
      },
      getRandomInt(min, max) {
        min = Math.ceil(min);
        max = Math.floor(max);
        return Math.floor(Math.random() * (max - min)) + min; //The maximum is exclusive and the minimum is inclusive
      }
    }
  };
</script>

<style lang="css">
  @import '../node_modules/bootstrap/dist/css/bootstrap.min.css';

  body {
    padding: 50px;
  }

  .chart {
    width: 100%;
    height: 400px;
    margin-bottom: 10px;
  }

  .chart-container {
    text-align: center;
    margin-top: 75px;
    margin-bottom: 15px;
  }

  .chart-list .chart-container:first-child {
    margin-top: 35px;
  }
</style>
