<template>
    <div>
        <el-badge :is-dot="isDot" class="item">
            <el-button type="primary" icon="bell" link @click="showSettingDrawer"
                :style="{ 'color': 'black', fontSize: '16px' }" />
        </el-badge>
        <!-- 消息通知 -->
        <el-drawer v-model="drawer" title="消息提示中心" :direction="direction" :before-close="handleClose">
            <el-tabs v-model="activeName" @tab-click="handleClick">
                <i class="el-icon-bell" slot="reference"></i>
                <el-tab-pane :label="'通知(' + count + ')'" name="first" style="color:#00aff0;">
                    <div v-for="(item, index) in userList" :key="index">
                        <el-button type="primary" icon="message" link @click="changeMessageState(item.ID)"
                            style="margin-top: 40px;padding: 0;">
                            {{ item.message_content }}
                        </el-button>
                        <el-divider style="margin: 0;padding: 0;"></el-divider>
                    </div>
                </el-tab-pane>
                <el-tab-pane label="消息(0)" name="third">未读</el-tab-pane>
                <el-tab-pane label="待办(0)" name="second">已读</el-tab-pane>
            </el-tabs>
        </el-drawer>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import {
    getMessage,
    setMessageState
} from '@/api/manage.js'

import { ElMessage } from 'element-plus'
import { useUserStore } from '@/pinia/modules/user'
import { useRouter } from 'vue-router'

const userStore = useUserStore()


defineOptions({
    name: 'mmsc',
})

onMounted(() => {
    getUserMessages()
})



const drawer = ref(false)
const direction = ref('rtl')
const isDot = ref(false)
const handleClose = () => {
    drawer.value = false
}
const showSettingDrawer = () => {
    drawer.value = true
}

const userList = ref([])

const count = ref(0)

const activeName = ref('first')

const getUserMessages = async () => {
    const table = await getMessage({ name: userStore.userInfo.userName })
    if (table.code === 0) {
        userList.value = table.data.list
        count.value = table.data.total
    }

    if (count.value !== 0) {
        isDot.value = true
    }
}
const router = useRouter()



const changeMessageState = async (id) => {
    const table = await setMessageState({ id: id })
    if (table.code != 0) {
        ElMessage({
            type: 'error',
            message: '更新信息失败',
            showClose: true
        })
        return
    }
    // getUserMessages()
    router.push({ name: 'abnormal' }).then(() => {
        window.location.reload()
    })
}



</script>

<style lang="scss" scoped>
.item {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-right: 13px;
}

.labelcolor {
    color: #00aff0;
    font-size: 16px;
    height: 50px;
    line-height: 50px;
    display: inline-block;
}

.el-tabs__item {
    height: 50px;
}

.notices {
    color: #c0c4cc;
    font-size: 14px;
    height: 100px;
    line-height: 100px;
    display: inline-block;
}
</style>