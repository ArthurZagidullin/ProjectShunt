<template>
  <v-layout
    justify-center
    align-center
  >
    <v-flex

      md12
    >
      <div class="text-xs-center">

      </div>
      <v-card>
        <v-card-title class="headline">Телеметрия датчиков</v-card-title>
        <v-card-text>
          <div class="small">
            <line-chart :chart-data="datacollection"></line-chart>
          </div>
        </v-card-text>
        <v-card-actions>
          <v-spacer />
        </v-card-actions>
      </v-card>
    </v-flex>
  </v-layout>
</template>

<script>
import Logo from '~/components/Logo.vue'
import VuetifyLogo from '~/components/VuetifyLogo.vue'
import LineChart from '~/components/LineChart'

export default {
  components: {
    Logo,
    LineChart,
    VuetifyLogo
  },
  data () {
    return {
      loaded: false,
      datacollection: null
    }
  },
  computed: {
    telemetry () {
      return this.$store.state.objects.telemetry
    }
  },
  watch: {
    telemetry(newData, oldData) {
      this.fillData(newData)
    }
  },
  mounted () {
    this.fillData([])
  },
  methods: {
    fillData (telemetryData) {
      var result = {
        labels: [],
        datasets: [],
      }

      var datasets = {
        label: 'Температура',
        backgroundColor: '#ff6384',
        data: [],
      }

      telemetryData.forEach((el)=>{
        if (!!el.data) {
          result.labels.push(el.time.getHours() + 'ч.' +el.time.getMinutes() + 'м.')
          datasets.data.push(el.data)
        }
      })

      result.datasets.push(datasets)
      console.log(result)

      this.datacollection = result
    },
    getRandomInt () {
      return Math.floor(Math.random() * (50 - 5 + 1)) + 5
    }
  }
}
</script>

<style>
  .small {
    max-width: 800px;
    margin:  50px auto;
  }
</style>
