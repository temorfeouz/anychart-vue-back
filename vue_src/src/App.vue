<template>
  <div class="container" id="app">
    <h1 class="text-center">
    Atmosphere info
    </h1>
    <section class="chart-list">
         <input placeholder="Metric period" id="period" value="10m" v-model="period">
      <input placeholder="Start period"  id="startTime"  value="" v-model="startTime">
      <input placeholder="End period"  id="endTime"  value="" v-model="endTime">
           <button class="btn btn-primary" @click="updateChartData">Update data
      </button>

      <tabs>
        <tab name="Street">
          <fieldset>
            <legend>Latest vals</legend>
            <h5 class="temp">Temp: {{ lastStreetTemp }} (~ {{ avgStreetTemp }}°C)</h5>
            <h5 class="humidity">Hum:  {{ lastStreetHum }}(~ {{ avgStreetHum }}%)</h5>
            <h5 >Gas:  <span v-html="lastStreetGas"></span> (~  <span v-html="avgStreetGas"></span> ) </h5>
            <h5 class="press">Press:{{ lastStreetPress }}(~ {{ avgStreetPress }})</h5>
          </fieldset>
          <section class="chart-container">
            <vue-anychart :options="TempStreet" ref="tempStreet"></vue-anychart>
          </section>
          <section class="chart-container">
            <vue-anychart :options="GasStreet" ref="gasStreet"></vue-anychart>
          </section>
          <section class="chart-container">
            <vue-anychart :options="HumStreet" ref="humStreet"></vue-anychart>
          </section>
          <section class="chart-container">
            <vue-anychart :options="PressStreet" ref="pressStreet"></vue-anychart>
          </section>
        </tab>
        <tab name="Room">
          <fieldset>
            <legend>Latest vals</legend>
            <h5 class="temp">Temp: {{ lastRoomTemp }} (~ {{ avgRoomTemp }}°C)</h5>
            <h5 class="humidity">Hum:  {{ lastRoomHum }}(~ {{ avgRoomHum }}%)</h5>
            <h5 >Gas:  <span v-html="lastRoomGas"></span> (~  <span v-html="avgRoomGas"></span> ) </h5>
            <h5 class="press">Press:{{ lastRoomPress }}(~ {{ avgRoomPress }})</h5>
          </fieldset>
          <section class="chart-container">
            <vue-anychart :options="TempRoom" ref="tempRoom"></vue-anychart>
          </section>
          <section class="chart-container">
            <vue-anychart :options="GasRoom" ref="gasRoom"></vue-anychart>
          </section>
          <section class="chart-container">
            <vue-anychart :options="HumRoom" ref="humRoom"></vue-anychart>
          </section>
          <section class="chart-container">
            <vue-anychart :options="PressRoom" ref="pressRoom"></vue-anychart>
          </section>
        </tab>
      </tabs>

    </section>
  </div>
</template>
<script>
/* eslint-disable */
function gasquality(val){
  if (val>0 &&val <=50){
    return "<span class='good'>Good</span>";
  }
  if (val>50 &&val <=100){
    return "<span class='avg'>Average</span>";
  }
  if (val>100 &&val <=150){
    return "<span class='litbad'>Little bad</span>";
  }
  if (val>150 &&val <=200){
    return "<span class='bad'>Bad</span>";
  }if (val>200 &&val <=300){
    return "<span class='worse'>Worse</span>";
  }if (val>300 &&val <=500){
    return "<span class='wevybad'>Very bad</span>";
  }

}

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
const timezoneOffset=2

function CurDate(){
 var ret= new Date()

  return new Date(ret.getFullYear(),ret.getMonth(), ret.getDate(), ret.getHours()+timezoneOffset, ret.getMinutes(), ret.getSeconds())
}
function calcAvg(elems){
  var ret=0;
  for (var i=0; i<elems.length;i++){
   ret+=elems[i][1]*1.0;
  }

  return parseFloat(ret/elems.length*1.0).toFixed(1)
}

function timeConverter(UNIX_timestamp, correctTZ){
  var a = new Date(UNIX_timestamp*1000 );
  var months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
  var year = a.getFullYear();
  var month = months[a.getMonth()];
  var date = a.getDate();
  if (correctTZ){
    var hour = a.getHours()+timezoneOffset;
  }else{
    var hour = a.getHours();
  }
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
        lastStreetTemp:"",
        lastStreetHum:"",
        lastStreetGas:"",
        lastStreetPress:"",
        avgStreetTemp:"",
        avgStreetHum:"",
        avgStreetGas:"",
        avgStreetPress:"",

        lastRoomTemp:"",
        lastRoomHum:"",
        lastRoomGas:"",
        lastRoomPress:"",
        avgRoomTemp:"",
        avgRoomHum:"",
        avgRoomGas:"",
        avgRoomPress:"",

        period:"10m",
        startTime:"",
        endTime:"",
        Anychart: Anychart,
        // areaOptions: data.AreaData,
        // pieOptions: data.PieData,
        // lineOptions: data.LineData,
        TempStreet: null,
        HumStreet: null,
        GasStreet:null,
        PressStreet:null,

        TempRoom: null,
        HumRoom: null,
        GasRoom:null,
        PressRoom:null,

        lineSeriesCount: 0,
        xAxisIsModified: false,
        pieDataIsModified: false
      }

      this.updateChartData()

      return ret
    },
    mounted() {
     var startdate = CurDate()
      startdate.setDate(startdate.getDate()-1)
      this.$data.startTime=ISODateString(startdate)
      this.$data.endTime=ISODateString(CurDate())


      // this.lineSeriesCount = this.$refs.lineChart.chart.getSeriesCount()
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

getlatestdata(){
  var std=CurDate()
  std.setHours(std.getHours()-12)
  var url ='http://192.168.1.253:8428/api/v1/query_range?start='+ISODateString(std)+'&end='+ISODateString(CurDate())+'&query={__name__=~"measurement.*"}'
console.log("latest data url ", url)
  axios.get(url).then(
    response =>  {
      response.data.data.result.forEach(v=>{
        if (v.metric.location=="street"){
          if (v.metric.__name__=="measurement_humidity"){
            var latest=v.values.length-1
            this.$data.avgStreetHum= calcAvg(v.values)
            this.$data.lastStreetHum= parseFloat(v.values[latest][1]).toFixed(1)+"% at "+timeConverter(v.values[latest][0])
          }
          if (v.metric.__name__=="measurement_temperature"){
            var latest=v.values.length-1
            this.$data.avgStreetTemp= calcAvg(v.values)
            this.$data.lastStreetTemp= parseFloat(v.values[latest][1]).toFixed(1)+"°C at "+timeConverter(v.values[latest][0])
          }
          if (v.metric.__name__=="measurement_gas"){
            var  latest=v.values.length-1
            this.$data.avgStreetGas= gasquality(calcAvg(v.values))
            this.$data.lastStreetGas=gasquality(v.values[latest][1])+" at "+timeConverter(v.values[latest][0])
          }
          if (v.metric.__name__=="measurement_preassure"){
            var   latest=v.values.length-1
            this.$data.avgStreetPress= calcAvg(v.values)
            this.$data.lastStreetPress=parseFloat(v.values[latest][1]).toFixed(1)+" at "+timeConverter(v.values[latest][0])
          }
        }
        if (v.metric.location=="home"){
          if (v.metric.__name__=="measurement_humidity"){
            var latest=v.values.length-1
            this.$data.avgRoomHum= calcAvg(v.values)
            this.$data.lastRoomHum= parseFloat(v.values[latest][1]).toFixed(1)+"% at "+timeConverter(v.values[latest][0])
          }
          if (v.metric.__name__=="measurement_temperature"){
            var latest=v.values.length-1
            this.$data.avgRoomTemp= calcAvg(v.values)
            this.$data.lastRoomTemp= parseFloat(v.values[latest][1]).toFixed(1)+"°C at "+timeConverter(v.values[latest][0])
          }
          if (v.metric.__name__=="measurement_gas"){
            var  latest=v.values.length-1
            this.$data.avgRoomGas= gasquality(calcAvg(v.values))
            this.$data.lastRoomGas=gasquality(v.values[latest][1])+" at "+timeConverter(v.values[latest][0])
          }
          if (v.metric.__name__=="measurement_preassure"){
            var   latest=v.values.length-1
            this.$data.avgRoomPress= calcAvg(v.values)
            this.$data.lastRoomPress=parseFloat(v.values[latest][1]).toFixed(1)+" at "+timeConverter(v.values[latest][0])
          }
        }
      })


    }).catch((error) => alert("ERRRR_"+error));

},

      updateChartData() {
        this.getlatestdata()
       var newPeriod=""
        var startdate = CurDate()
        startdate.setDate(startdate.getDate()-1)
        var startTime=ISODateString(startdate)
        var endTime=ISODateString(CurDate())
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

            this.$data.TempStreet= {
              'chart': {
                'title': 'Temp',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [ {
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
            this.$data.HumStreet= {
              'chart': {
                'title': 'Humidity',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [ {
                  'seriesType': 'spline',
                  'name': 'Humidity',
                  'normal': {'stroke': {'color': 'blue', 'thickness': 2.5}},
                  'data': [
                  ]
                }
                ],
                'type': 'column'
              }
            }

            this.$data.GasStreet= {
              'chart': {
                'title': 'Gas',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [  {
                  'seriesType': 'spline',
                  'name': 'Gas',
                  'normal': {'stroke': {'color': 'black', 'thickness': 2.5}},
                  'data': [
                  ]
                }
                ],
                'type': 'column'
              }
            }
            this.$data.PressStreet= {
              'chart': {
                'title': 'Pressure',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [{
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

            ///////////////////////
            this.$data.TempRoom= {
              'chart': {
                'title': 'Temp',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [ {
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
            this.$data.HumRoom= {
              'chart': {
                'title': 'Humidity',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [ {
                  'seriesType': 'spline',
                  'name': 'Humidity',
                  'normal': {'stroke': {'color': 'blue', 'thickness': 2.5}},
                  'data': [
                  ]
                }
                ],
                'type': 'column'
              }
            }

            this.$data.GasRoom= {
              'chart': {
                'title': 'Gas',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [  {
                  'seriesType': 'spline',
                  'name': 'Gas',
                  'normal': {'stroke': {'color': 'black', 'thickness': 2.5}},
                  'data': [
                  ]
                }
                ],
                'type': 'column'
              }
            }
            this.$data.PressRoom= {
              'chart': {
                'title': 'Pressure',
                'animation': true,
                'tooltip': {'displayMode': 'union'},
                'interactivity': {'hoverMode': 'by-x'},
                'series': [{
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
                if (v.metric.location=="street"){
                  if (v.metric.__name__=="measurement_humidity"){
                    for ( var i=0;i<v.values.length;i++){
                      if (v.values[i][1]==0){continue}
                      this.$data.HumStreet.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                  if (v.metric.__name__=="measurement_temperature"){
                    for ( var i=0;i<v.values.length;i++){
                      if (v.values[i][1]< -60){continue}
                      this.$data.TempStreet.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                  if (v.metric.__name__=="measurement_gas"){
                    for ( var i=0;i<v.values.length;i++){
                      // if (v.values[i][1]==0){continue}
                      this.$data.GasStreet.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                  if (v.metric.__name__=="measurement_preassure"){
                    for ( var i=0;i<v.values.length;i++){
                      if (v.values[i][1]==0){continue}
                      this.$data.PressStreet.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                }
                if (v.metric.location=="home"){
                  if (v.metric.__name__=="measurement_humidity"){
                    for ( var i=0;i<v.values.length;i++){
                      if (v.values[i][1]==0){continue}
                      this.$data.HumRoom.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                  if (v.metric.__name__=="measurement_temperature"){
                    for ( var i=0;i<v.values.length;i++){
                      if (v.values[i][1]< -60){continue}
                      this.$data.TempRoom.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                  if (v.metric.__name__=="measurement_gas"){
                    for ( var i=0;i<v.values.length;i++){
                      // if (v.values[i][1]==0){continue}
                      this.$data.GasRoom.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                  if (v.metric.__name__=="measurement_preassure"){
                    for ( var i=0;i<v.values.length;i++){
                      if (v.values[i][1]==0){continue}
                      this.$data.PressRoom.chart.series[0].data.push({'x': timeConverter(v.values[i][0]), 'value': parseFloat(v.values[i][1]).toFixed(0)})
                    }
                  }
                }
            }
                          )


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
#period{
  width: 32pt;
}
#startTime,#endTime{
  width: 129pt;
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
  fieldset h5 {
    min-width: 0;
    padding: 0;
    margin: 0;
    border: 0;
    font-size: 100%;
  }
  .humidity{
    color:blue;
  }
  .temp{
    color:red;
  }
  .gas{
    color: black;
  }
  .press{
    color: green;
  }
</style>
