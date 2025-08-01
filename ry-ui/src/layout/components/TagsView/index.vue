<template>
	<div id="tags-view-container" class="tags-view-container">
		<scroll-pane
			ref="scrollPaneRef"
			class="tags-view-wrapper"
			@scroll="handleScroll"
		>
			<router-link
				v-for="tag in visitedViews"
				:key="tag.path"
				:data-path="tag.path"
				:class="isActive(tag) ? 'active' : ''"
				:to="{
					path: tag.path,
					query: tag.query,
					fullPath: tag.fullPath,
				} as any"
				class="tags-view-item"
				:style="activeStyle(tag)"
				@click.middle="!isAffix(tag) ? closeSelectedTag(tag) : ''"
				@contextmenu.prevent="openMenu(tag, $event)"
			>
				{{ tag.title }}
				<span
					v-if="!isAffix(tag)"
					@click.prevent.stop="closeSelectedTag(tag)"
				>
					<close
						class="close"
						style="width: 1em; height: 1em; vertical-align: middle"
					/>
				</span>
			</router-link>
		</scroll-pane>
		<ul
			v-show="visible"
			:style="{ left: left + 'px', top: top + 'px' }"
			class="contextmenu"
		>
			<li @click="refreshSelectedTag(selectedTag!)">
				<refresh-right style="width: 1em; height: 1em" /> 刷新页面
			</li>
			<li
				v-if="!isAffix(selectedTag!)"
				@click="closeSelectedTag(selectedTag!)"
			>
				<close style="width: 1em; height: 1em" /> 关闭当前
			</li>
			<li @click="closeOthersTags">
				<circle-close style="width: 1em; height: 1em" /> 关闭其他
			</li>
			<li v-if="!isFirstView()" @click="closeLeftTags">
				<back style="width: 1em; height: 1em" /> 关闭左侧
			</li>
			<li v-if="!isLastView()" @click="closeRightTags">
				<right style="width: 1em; height: 1em" /> 关闭右侧
			</li>
			<li @click="closeAllTags(selectedTag!)">
				<circle-close style="width: 1em; height: 1em" /> 全部关闭
			</li>
		</ul>
	</div>
</template>

<script setup lang="ts">
// prettier-ignore
import { ref, computed, nextTick, watch, onMounted, getCurrentInstance, } from "vue";
import { RouteRecordRaw, useRoute, useRouter } from "vue-router";
import ScrollPane from "./ScrollPane.vue";
import { getNormalPath, useSafeInstance } from "@/utils/ruoyi";
import useTagsViewStore, { VisitedView } from "@/store/modules/tagsView";
import useSettingsStore from "@/store/modules/settings";
import usePermissionStore from "@/store/modules/permission";

const visible = ref(false);
const top = ref(0);
const left = ref(0);
const selectedTag = ref<VisitedView | null>(null);
const affixTags = ref<VisitedView[]>([]);
const scrollPaneRef = ref<InstanceType<typeof ScrollPane> | null>(null);

const proxy = useSafeInstance();
const route = useRoute();
const router = useRouter();

const visitedViews = computed(() => useTagsViewStore().visitedViews);
const routes = computed(() => usePermissionStore().routes);
const theme = computed(() => useSettingsStore().theme);

watch(route, () => {
	addTags();
	moveToCurrentTag();
});
watch(visible, (value) => {
	if (value) {
		document.body.addEventListener("click", closeMenu);
	} else {
		document.body.removeEventListener("click", closeMenu);
	}
});
onMounted(() => {
	initTags();
	addTags();
});

function isActive(r: VisitedView) {
	return r.path === route.path;
}
function activeStyle(tag: VisitedView) {
	if (!isActive(tag)) return {};
	return {
		"background-color": theme.value,
		"border-color": theme.value,
	};
}
const isAffix = (tag: VisitedView): boolean => {
	return !!tag && !!tag.meta && tag.meta.affix!!;
}
function isFirstView() {
	try {
		return (
			selectedTag.value!.fullPath === visitedViews.value[1].fullPath ||
			selectedTag.value!.fullPath === "/index"
		);
	} catch (err) {
		return false;
	}
}
function isLastView() {
	try {
		return (
			selectedTag.value!.fullPath ===
			visitedViews.value[visitedViews.value.length - 1].fullPath
		);
	} catch (err) {
		return false;
	}
}
function filterAffixTags(routes: RouteRecordRaw[], basePath = "") {
	let tags: VisitedView[] = [];
	routes.forEach((route) => {
		if (route.meta && route.meta.affix) {
			const tagPath = getNormalPath(basePath + "/" + route.path);
			tags.push({
				fullPath: tagPath,
				path: tagPath,
				name: route.name,
				meta: { ...route.meta },
				title: (route.meta?.title as string) || "",
			});
		}
		if (route.children) {
			const tempTags = filterAffixTags(route.children, route.path);
			if (tempTags.length >= 1) {
				tags = [...tags, ...tempTags];
			}
		}
	});
	return tags;
}
function initTags() {
	const res = filterAffixTags(routes.value);
	affixTags.value = res;
	for (const tag of res) {
		// Must have tag name
		if (tag.name) {
			useTagsViewStore().addVisitedView(tag as any);
		}
	}
}
function addTags() {
	const { name } = route;
	if (name) {
		useTagsViewStore().addView(route);
	}
	return false;
}
function moveToCurrentTag() {
	nextTick(() => {
		for (const r of visitedViews.value) {
			if (r.path === route.path && scrollPaneRef.value) {
				scrollPaneRef.value.moveToTarget(r);
				if (r.fullPath !== route.fullPath) {
					useTagsViewStore().updateVisitedView(route);
				}
			}
		}
	});
}
function refreshSelectedTag(view: VisitedView) {
	proxy.$tab.refreshPage(view);
}
const closeSelectedTag = async (view: VisitedView): Promise<void> => {
	const { visitedViews } = await proxy.$tab.closePage(view);
	if (isActive(view)) {
		toLastView(visitedViews, view);
	}
};

function closeRightTags() {
	proxy.$tab.closeRightPage(selectedTag.value).then((visitedViews: any[]) => {
		if (
			!visitedViews.find(
				(i: { fullPath: string }) => i.fullPath === route.fullPath
			)
		) {
			toLastView(visitedViews);
		}
	});
}

const closeLeftTags = async (): Promise<void> => {
	const visitedViews = await proxy.$tab.closeLeftPage(selectedTag.value!);
	if (
		!visitedViews.find(
			(i: { fullPath: string }) => i.fullPath === route.fullPath
		)
	) {
		toLastView(visitedViews);
	}
};

async function closeOthersTags() {
	await router.push(selectedTag.value!.fullPath!);
	await proxy.$tab.closeOtherPage(selectedTag.value!);
	moveToCurrentTag();
}
async function closeAllTags(view: VisitedView) {
	const { visitedViews } = await proxy.$tab.closeAllPage(view);
	if (affixTags.value.some((tag) => tag.path === route.path)) return;
	toLastView(visitedViews, view);
}
function toLastView(visitedViews: VisitedView[], view?: VisitedView) {
	const latestView = visitedViews.slice(-1)[0];
	if (latestView) {
		router.push(latestView.fullPath!);
	} else {
		// now the default is to redirect to the home page if there is no tags-view,
		// you can adjust it according to your needs.
		if (view!.name === "Dashboard") {
			// to reload home page
			router.replace({ path: "/redirect" + view!.fullPath });
		} else {
			router.push("/");
		}
	}
}
function openMenu(tag: VisitedView, e: MouseEvent) {
	const menuMinWidth = 105;
	const offsetLeft = proxy.$el.getBoundingClientRect().left; // container margin left
	const offsetWidth = proxy.$el.offsetWidth; // container width
	const maxLeft = offsetWidth - menuMinWidth; // left boundary
	const l = e.clientX - offsetLeft + 15; // 15: margin right

	if (l > maxLeft) {
		left.value = maxLeft;
	} else {
		left.value = l;
	}

	top.value = e.clientY;
	visible.value = true;
	selectedTag.value = tag;
}
function closeMenu() {
	visible.value = false;
}
function handleScroll() {
	closeMenu();
}
</script>

<style lang="scss" scoped>
.tags-view-container {
	height: 34px;
	width: 100%;
	background: #fff;
	border-bottom: 1px solid #d8dce5;
	box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.12), 0 0 3px 0 rgba(0, 0, 0, 0.04);
	.tags-view-wrapper {
		.tags-view-item {
			display: inline-block;
			position: relative;
			cursor: pointer;
			height: 26px;
			line-height: 26px;
			border: 1px solid #d8dce5;
			color: #495060;
			background: #fff;
			padding: 0 8px;
			font-size: 12px;
			margin-left: 5px;
			margin-top: 4px;
			&:first-of-type {
				margin-left: 15px;
			}
			&:last-of-type {
				margin-right: 15px;
			}
			&.active {
				background-color: #42b983;
				color: #fff;
				border-color: #42b983;
				&::before {
					content: "";
					background: #fff;
					display: inline-block;
					width: 8px;
					height: 8px;
					border-radius: 50%;
					position: relative;
					margin-right: 2px;
				}
			}
		}
	}
	.contextmenu {
		margin: 0;
		background: #fff;
		z-index: 3000;
		position: absolute;
		list-style-type: none;
		padding: 5px 0;
		border-radius: 4px;
		font-size: 12px;
		font-weight: 400;
		color: #333;
		box-shadow: 2px 2px 3px 0 rgba(0, 0, 0, 0.3);
		li {
			margin: 0;
			padding: 7px 16px;
			cursor: pointer;
			&:hover {
				background: #eee;
			}
		}
	}
}
</style>

<style lang="scss">
//reset element css of close
.tags-view-wrapper {
	.tags-view-item {
		.close {
			width: 16px;
			height: 16px;
			vertical-align: 2px;
			border-radius: 50%;
			text-align: center;
			transition: all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
			transform-origin: 100% 50%;
			&:before {
				transform: scale(0.6);
				display: inline-block;
				vertical-align: -3px;
			}
			&:hover {
				background-color: #b4bccc;
				color: #fff;
				width: 12px !important;
				height: 12px !important;
			}
		}
	}
}
</style>
