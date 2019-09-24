export function parseTime(time, cFormat) {
  if (arguments.length === 0) {
    return null;
  }
  const format = cFormat || "{y}-{m}-{d} {h}:{i}:{s}";
  let date;
  if (typeof time === "object") {
    date = time;
  } else {
    if (("" + time).length === 10) time = parseInt(time) * 1000;
    date = new Date(time);
  }
  const formatObj = {
    y: date.getFullYear(),
    m: date.getMonth() + 1,
    d: date.getDate(),
    h: date.getHours(),
    i: date.getMinutes(),
    s: date.getSeconds(),
    a: date.getDay()
  };
  const time_str = format.replace(/{(y|m|d|h|i|s|a)+}/g, (result, key) => {
    let value = formatObj[key];
    if (key === "a")
      return ["一", "二", "三", "四", "五", "六", "日"][value - 1];
    if (result.length > 0 && value < 10) {
      value = "0" + value;
    }
    return value || 0;
  });
  return time_str;
}

export function formatTime(time, option) {
  time = +time * 1000;
  const d = new Date(time);
  const now = Date.now();

  const diff = (now - d) / 1000;

  if (diff < 30) {
    return "刚刚";
  } else if (diff < 3600) {
    // less 1 hour
    return Math.ceil(diff / 60) + "分钟前";
  } else if (diff < 3600 * 24) {
    return Math.ceil(diff / 3600) + "小时前";
  } else if (diff < 3600 * 24 * 2) {
    return "1天前";
  }
  if (option) {
    return parseTime(time, option);
  } else {
    return (
      d.getMonth() +
      1 +
      "月" +
      d.getDate() +
      "日" +
      d.getHours() +
      "时" +
      d.getMinutes() +
      "分"
    );
  }
}

export function int2ip(i) {
  let ip = "";
  if (i <= 0) {
    return "0.0.0.0";
  }
  var ip3 = (i << 0) >>> 24;
  var ip2 = (i << 8) >>> 24;
  var ip1 = (i << 16) >>> 24;
  var ip0 = (i << 24) >>> 24;
  ip += ip3 + "." + ip2 + "." + ip1 + "." + ip0;
  return ip;
}

export function int2ip_reverse(i) {
  let ip = "";
  if (i <= 0) {
    return "0.0.0.0";
  }
  var ip3 = (i << 0) >>> 24;
  var ip2 = (i << 8) >>> 24;
  var ip1 = (i << 16) >>> 24;
  var ip0 = (i << 24) >>> 24;
  ip += ip0 + "." + ip1 + "." + ip2 + "." + ip3;
  return ip;
}

export function formatKbps(value) {
  //value单位为Kbit
  if (value > 1 << 20) {
    //GB
    return (value / (1 << 20)).toFixed(2) + " Gbps";
  }
  if (value > 1 << 10) {
    return (value / (1 << 10)).toFixed(2) + " Mbps";
  }
  return value + " Kbps";
}

export function ip2int(ip) {
  var num = 0;
  ip = ip.split(".");
  num =
    Number(ip[0]) * 256 * 256 * 256 +
    Number(ip[1]) * 256 * 256 +
    Number(ip[2]) * 256 +
    Number(ip[3]);
  num = num >>> 0;
  return num;
}
