<template>
	<div class="navbar">
		<hamburger
			id="hamburger-container"
			:is-active="appStore.sidebar.opened"
			class="hamburger-container"
			@toggleClick="toggleSideBar"
		/>
		<breadcrumb
			id="breadcrumb-container"
			class="breadcrumb-container"
			v-if="!settingsStore.topNav"
		/>
		<top-nav
			id="topmenu-container"
			class="topmenu-container"
			v-if="settingsStore.topNav"
		/>

		<div class="right-menu">
			<template v-if="appStore.device !== 'mobile'">
                <span style="font-size: 12px;margin-right: 20px;color: #696969;">上次登录：{{ lastLogin }}</span>
				<header-search id="header-search" class="right-menu-item" />
				<el-tooltip content="源码地址" effect="dark" placement="bottom">
					<ruo-yi-git
						id="ruoyi-git"
						class="right-menu-item hover-effect"
					/>
				</el-tooltip>

				<el-tooltip content="文档地址" effect="dark" placement="bottom">
					<ruo-yi-doc
						id="ruoyi-doc"
						class="right-menu-item hover-effect"
					/>
				</el-tooltip>

				<screen-full
					id="screenFull"
					class="right-menu-item hover-effect"
				/>

				<el-tooltip content="布局大小" effect="dark" placement="bottom">
					<size-select
						id="size-select"
						class="right-menu-item hover-effect"
                         style="margin-top: 13px"
					/>
				</el-tooltip>
			</template>
			<div class="avatar-container">
				<el-dropdown
					@command="handleCommand"
					class="right-menu-item hover-effect"
					trigger="click"
				>
					<div class="avatar-wrapper">
						<img :src="userStore.avatar" class="user-avatar" />
						<el-icon><caret-bottom /></el-icon>
					</div>
					<template #dropdown>
						<el-dropdown-menu>
							<router-link to="/user/profile">
								<el-dropdown-item>个人中心</el-dropdown-item>
							</router-link>
							<el-dropdown-item command="setLayout">
								<span>布局设置</span>
							</el-dropdown-item>
							<el-dropdown-item divided command="logout">
								<span>退出登录</span>
							</el-dropdown-item>
						</el-dropdown-menu>
					</template>
				</el-dropdown>
			</div>
		</div>
	</div>
</template>

<script setup>
import router from "@/router";
import { ref, onMounted, watch } from "vue";
import { ElMessageBox } from "element-plus";
import Breadcrumb from "@/components/Breadcrumb/index.vue";
import TopNav from "@/components/TopNav/index.vue";
import Hamburger from "@/components/Hamburger/index.vue";
import ScreenFull from "@/components/ScreenFull/index.vue";
import SizeSelect from "@/components/SizeSelect/index.vue";
import HeaderSearch from "@/components/HeaderSearch/index.vue";
import RuoYiGit from "@/components/RuoYi/Git/index.vue";
import RuoYiDoc from "@/components/RuoYi/Doc/index.vue";
import useAppStore from "@/store/modules/app";
import useUserStore from "@/store/modules/user";
import useSettingsStore from "@/store/modules/settings";
import request from "@/utils/request";

const appStore = useAppStore();
const userStore = useUserStore();
const settingsStore = useSettingsStore();
const lastLogin = ref("");

function toggleSideBar() {
	appStore.toggleSideBar();
}
function handleCommand(command) {
	switch (command) {
		case "setLayout":
			setLayout();
			break;
		case "logout":
			logout();
			break;
		default:
			break;
	}
}
function logout() {
	// prettier-ignore
	ElMessageBox.confirm("确定注销并退出系统吗？", "提示", {
		confirmButtonText: "确定",
		cancelButtonText: "取消",
		type: "warning",
	})
    .then(() => {
        userStore.logOut().then(() => {
            window.location.href = "/index";
        });
    })
    .catch(() => {});
}
const emits = defineEmits(["setLayout"]);
function setLayout() {
	emits("setLayout");
}

const lastLoginTime = async () => {
    await request({
        url: "/monitor/logininfor/lastLogin",
        method: "get",
    }).then(response => {
        if (response.code === 200) {
            lastLogin.value = response.data;
        }
    });
};

onMounted(() => {
    //lastLoginTime();
});

watch(() => router.currentRoute.value.path,(newValue, oldValue) => {
        console.log("路由变化", newValue, "刷新距离上次登录时间");
        //lastLoginTime()
    },
    { immediate: true } // 初始化之后立即调用。
);

// setInterval(() => {
//    lastLoginTime();
// }, 1000 * 60);

</script>

<style lang="scss" scoped>
.navbar {
	height: 50px;
	overflow: hidden;
	position: relative;
	background: #fff;
	box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

	.hamburger-container {
		line-height: 46px;
		height: 100%;
		float: left;
		cursor: pointer;
		transition: background 0.3s;
		-webkit-tap-highlight-color: transparent;

		&:hover {
			background: rgba(0, 0, 0, 0.025);
		}
	}

	.breadcrumb-container {
		float: left;
	}

	.topmenu-container {
		position: absolute;
		left: 50px;
	}

	.errLog-container {
		display: inline-block;
		vertical-align: top;
	}

	.right-menu {
		float: right;
		height: 100%;
		line-height: 50px;
		display: flex;

		&:focus {
			outline: none;
		}

		.right-menu-item {
			display: inline-block;
			padding: 0 8px;
			height: 100%;
			font-size: 18px;
			color: #5a5e66;
			vertical-align: text-bottom;

			&.hover-effect {
				cursor: pointer;
				transition: background 0.3s;

				&:hover {
					background: rgba(0, 0, 0, 0.025);
				}
			}
		}

		.avatar-container {
			margin-right: 40px;

			.avatar-wrapper {
				margin-top: 5px;
				position: relative;

				.user-avatar {
					cursor: pointer;
					width: 40px;
					height: 40px;
					border-radius: 10px;
				}

				i {
					cursor: pointer;
					position: absolute;
					right: -20px;
					top: 25px;
					font-size: 12px;
				}
			}
		}
	}
}
</style>
