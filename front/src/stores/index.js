import { createStore } from "vuex";
import createPersistedState from "vuex-persistedstate"

const store = createStore({
  plugins: [createPersistedState({})],
  state: {
    Token: "",
    userMsg: {},
    tags: [],
  },
  mutations: {
    setToken(state, token){
      state.Token = token
    },
    setUserMsg(state, msg){
      state.userMsg = msg
    },
    setTags(state, tags){
      state.tags = tags
    },
    setTeaId(state, id){
      state.userMsg.Yb_studentid = id
    }
  }
})

export default store