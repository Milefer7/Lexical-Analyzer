<!-- src/views/PopupView.vue -->
<template>
  <div>
    <component :is="currentComponent"></component>
  </div>
</template>

<script>
import CharTable from "../components/CharTable.vue"
import TokenClassification from "../components/TokenClassification.vue"
import KeywordTable from "../components/KeywordTable.vue"
import DFA from "../components/DFA.vue"
import ParsingProcess from "../components/ParsingProcess.vue"
import DelimeterTable from "@/components/DelimeterTable.vue"
import SyntaxRules from "@/components/SyntaxRules.vue"

export default {
  computed: {
    currentComponent() {
      const component = this.$route.query.component
      const components = {
        CharTable,
        TokenClassification,
        KeywordTable,
        DFA,
        ParsingProcess,
        DelimeterTable,
        SyntaxRules,
      }
      return components[component] || null
    },
  },
}
</script>
