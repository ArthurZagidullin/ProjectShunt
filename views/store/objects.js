import Vue from 'vue'

export const state = () => ({
  telemetry: [],
})


export const mutations = {
  NewData(state, data) {
    state.telemetry.push(data)
  },
}
export const actions = {
  NewData({commit}, data) {
    console.log(data)
    if (!!data['data']) {
      var dt = new Date(data['timestamp'])
      commit('NewData', {data: data['data'], time: dt})
    }
  },
}
