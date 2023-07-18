<script lang="ts" setup>
import { useRoute } from 'vue-router';
import { routes } from '../router'
import { computed } from 'vue';

const menus = routes.filter(it => it.type == 'menu')

const route = useRoute()

const activePath = computed(() => {
    return route.path == "/" ? menus[0].path : route.path
})
</script>

<template>
    <div class="common-layout">
        <el-container>
            <el-aside style="width: 180px;">
                <el-row>
                    <el-col>
                        <h5 class="h5-title">有声小说工厂</h5>
                        <el-menu :router="true" active-text-color="#ffd04b" background-color="#545c64" text-color="#fff"
                            class="el-menu-vertical" :default-active=activePath>
                            <el-menu-item v-for="menu in menus" :index="menu.path">
                                <el-icon>
                                    <component :is="menu.icon" />
                                </el-icon>
                                <span>{{ menu.name }}</span>
                            </el-menu-item>
                        </el-menu>
                    </el-col>
                </el-row>
            </el-aside>
            <el-main>
                <router-view></router-view>
            </el-main>
        </el-container>
    </div>
</template>

<style scoped>
.h5-title {
    text-align: left;
    padding-left: 20px;
    font-size: large;
    color: #e7ac08;
}

.el-aside {
    background-color: #545c64;
}

.el-menu-vertical {
    border: 0;
}

.el-container {
    height: 100vh;
}
</style>
