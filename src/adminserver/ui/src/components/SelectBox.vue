<template>
  <Listbox v-model="selectedValue">
    <div class="relative w-full">
      <ListboxButton
          class="w-full inline-flex items-center py-2 px-3 border rounded-md border-gray-300 sm:text-sm focus:outline-indigo-500">
        <span class="flex-grow text-left">{{ selectedFormat }}</span>
        <ChevronUpDownIcon class="w-5 h-5 text-gray-400" aria-hidden="true"/>
      </ListboxButton>
      <Transition leave-active-class="transition duration-100 ease-in"
                  leave-from-class="opacity-100"
                  leave-to-class="opacity-0">
        <ListboxOptions
            class="mt-1 absolute z-50 grid grid-cols-1 max-h-60 w-full overflow-y-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
          <ListboxOption v-for="opt in options" :key="opt" :value="value(opt)" v-slot="{ selected }"
                         class="inline-flex items-center select-none hover:bg-indigo-200 hover:bg-opacity-50 hover:text-indigo-600">
            <CheckIcon class="w-5 h-5 text-indigo-500 mx-3" :class="{'opacity-0': !selected}"/>
            <span class="py-2 text-left" :class="{
              'font-medium': selected
            }">{{ format(opt) }}</span>
          </ListboxOption>
          <div v-if="options.length === 0" class="text-center text-gray-400 py-3">{{ t('no-option') }}</div>
        </ListboxOptions>
      </Transition>
    </div>
  </Listbox>
</template>

<script lang="ts" setup>
import {CheckIcon, ChevronUpDownIcon} from '@heroicons/vue/20/solid'

const prop = withDefaults(defineProps<{
  modelValue: any
  options: any[]
  value: (v: any) => any
  format: (v: any) => any
}>(), {
  format: (v: any) => v,
  key: (v: any) => v,
  value: (v: any) => v
})

const emit = defineEmits(['update:modelValue'])

const selectedValue = useVModel(prop, 'modelValue', emit)

const {t} = useI18n({
  inheritLocale: true
})

const selectedFormat = computed(()=>{
  if (!selectedValue) {
    return ''
  }
  for (let i in prop.options) {
    if(prop.value(prop.options[i]) === selectedValue.value){
      return prop.format(prop.options[i])
    }
  }
  return ''
})
</script>

<i18n lang="yaml">
en:
  no-option: 'No options'
zh:
  no-option: '没有待选项'
</i18n>