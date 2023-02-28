export default{
	state:{
		isCollapse:false,//用于控制菜单展开
	},
	mutations:{
		// 修改菜单展开
		collapaseMenu(state){
			state.isCollapse = !state.isCollapse
		}
	}
}
