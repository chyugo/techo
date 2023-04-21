<template>
	<div>
		<el-container>
			<el-aside width="auto"   style="background-color: rgb(84, 92, 100);">
				<common-aside></common-aside>
			</el-aside>
			<el-container style="position: relative;">
				<el-header style="position: sticky;top: 0px;z-index: 999;">
					<common-header></common-header>
				</el-header>
				<el-main>
					<router-view></router-view>
				</el-main>
				
				<common-footer></common-footer>
			</el-container>
		</el-container>
	</div>
</template>

<script>
	import CommonAside from '../components/CommonAside.vue'
	import CommonHeader from '../components/CommonHeader.vue'
	import CommonFooter from '../components/CommonFooter.vue'
	export default {
		data() {
			return {
				
			}
		},
		components:{
			CommonAside,
			CommonHeader,
			CommonFooter
		}
	}
</script>

<style>
	.el-header{
		padding: 0;
	}
</style>
