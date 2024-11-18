module.exports = {
  displayName: "frontend",
  preset: "../jest.preset.js",
  transform: {
    "^.+.vue$": "@vue/vue3-jest",
    ".+.(css|styl|less|sass|scss|svg|png|jpg|ttf|woff|woff2)$":
      "jest-transform-stub",
    "^.+.tsx?$": "ts-jest",
  },
  moduleFileExtensions: ["ts", "tsx", "vue", "js", "json"],
  coverageDirectory: "../coverage/./frontend",
  snapshotSerializers: ["jest-serializer-vue"],
  globals: {
    "ts-jest": {
      tsconfig: "./frontend/tsconfig.spec.json",
    },
    "vue-jest": {
      tsConfig: "./frontend/tsconfig.spec.json",
    },
  },
};
