<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="ltacker.jupiter.laugh" type="hoho" :fields="[ ['creator', 1, 'string'] , ['text', 2, 'string'] ]" />
		<sp-type-form path="ltacker.jupiter.laugh" type="hihi" :fields="[ ['creator', 1, 'string'] , ['text', 2, 'string'] ]" />
		<sp-type-form path="ltacker.jupiter.laugh" type="haha" :fields="[ ['creator', 1, 'string'] , ['text', 2, 'string'] ]" />
		<sp-type-form path="ltacker.jupiter.laugh" type="hohosent" :fields="[ ['creator', 1, 'string'] , ['text', 2, 'string'] ]" />
		<sp-type-form path="ltacker.jupiter.laugh" type="hihisent" :fields="[ ['creator', 1, 'string'] , ['text', 2, 'string'] ]" />
		<sp-type-form path="ltacker.jupiter.laugh" type="hahasent" :fields="[ ['creator', 1, 'string'] , ['text', 2, 'string'] ]" />
		<sp-type-form path="ltacker.jupiter.ibcchat" type="message" :fields="[ ['creator', 1, 'string'] , ['text', 2, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
