<template>
  <div class="text-2xl text-gray-900 font-bold mb-4">{{ $t('overview') }}</div>
  <div v-if="infos.length > 0"
       class="grid gap-y-4 grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 2xl:grid-cols-8 justify-items-center placeholder:py-2">
    <ServerCard class="cursor-pointer" v-for="info in infos" :info="info"
                @click="openDialog(info.serverId)"></ServerCard>
  </div>
  <div v-else class="w-full my-5 text-center text-gray-600 text-lg font-medium">
    {{ $t('no-servers') }}
  </div>
  <div class="mb-4 mt-8">
    <!-- capacity card -->
    <CapCard v-if="capInfo.total > 0" class="w-[32%]" :cap-info="capInfo"/>
  </div>
  <div class="mt-8 text-2xl text-gray-900 font-bold mb-4">{{ $t('monitor') }}</div>
  <UsageLine h="h-56" class="mb-4" :type="$cst.statTypeCpu" :tl="cpuTl"/>
  <UsageLine h="h-56" :type="$cst.statTypeMem" :tl="memTl"/>
  <!-- Migration dialog -->
  <ModalTemplate title="Join or Leave" v-model="openMigrateDialog">
    <template #panel>
      <div class="py-6 px-8 grid-cols-1 grid gap-y-2">
        <button class="btn-pri" @click="clusterCmd('leave')">
          <font-awesome-icon v-if="inRequesting" class="animate-spin mr-2" icon="spinner"/>
          Leave cluster
        </button>
        <button class="btn-revert" @click="openMigrateDialog = false">Close</button>
      </div>
    </template>
  </ModalTemplate>
</template>

<script setup lang="ts">
const infos = ref<ServerInfo[]>([])
const capInfo = ref<DiskInfo>({used: 0, total: 0, free: 0})
const cpuTl = ref<Record<string, TimeStat[]>>({})
const memTl = ref<Record<string, TimeStat[]>>({})
const store = useStore()
const openMigrateDialog = ref(false)
const inRequesting = ref(false)
let migrateServId = ""
const {t} = useI18n({inheritLocale: true})

function updateInfo(state: any) {
    let infoList: ServerInfo[] = []
    let cap = {used: 0, total: 0, free: 0}
    let stats = state.serverStat.dataServer
    for (let k in stats) {
        let v = stats[k]
        infoList.push(v)
        cap.used += v.sysInfo.diskInfo.used
        cap.total += v.sysInfo.diskInfo.total
        cap.free += v.sysInfo.diskInfo.free
    }
    infos.value = infoList
    capInfo.value = cap
}

function openDialog(servId: string) {
    migrateServId = servId
    openMigrateDialog.value = true
}

async function clusterCmd(cmd: string) {
    try {
        inRequesting.value = true
        if (cmd == 'join') {
            await api.objects.join(migrateServId)
        } else if (cmd == 'leave') {
            await api.objects.leave(migrateServId)
        }
        openMigrateDialog.value = false
        useToast().success(t('req-success'))
    } catch (err: any) {
        useToast().error(err.message)
    } finally {
        inRequesting.value = false
    }
}

function getTl() {
    api.serverStat.timeline(pkg.cst.dataServerNo, pkg.cst.statTypeCpu)
        .then(res => {
            cpuTl.value = res
        })
    api.serverStat.timeline(pkg.cst.dataServerNo, pkg.cst.statTypeMem)
        .then(res => {
            memTl.value = res
        })
}

store.$subscribe((mutation, state) => {
    updateInfo(state)
})

onBeforeMount(() => {
    updateInfo(store)
    pkg.utils.invokeInterval(() => {
        getTl()
    }, 1000 * 60 * 60)
})
</script>

<style scoped>

</style>

<route lang="json">
{
  "meta": {
    "title": "data-server"
  }
}
</route>