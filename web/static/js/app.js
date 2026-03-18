// Графики
let cpuChart, memChart;

function initCharts() {
    const cpuCtx = document.getElementById('cpu-chart').getContext('2d');
    cpuChart = new Chart(cpuCtx, {
        type: 'line',
        data: {
            labels: [],
            datasets: [{
                label: 'CPU Usage %',
                data: [],
                borderColor: '#3b82f6',
                backgroundColor: 'rgba(59, 130, 246, 0.1)',
                tension: 0.3,
                fill: true
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            scales: { y: { beginAtZero: true, max: 100 } },
            plugins: { legend: { display: false } }
        }
    });

    const memCtx = document.getElementById('mem-chart').getContext('2d');
    memChart = new Chart(memCtx, {
        type: 'line',
        data: {
            labels: [],
            datasets: [{
                label: 'Memory Usage %',
                data: [],
                borderColor: '#10b981',
                backgroundColor: 'rgba(16, 185, 129, 0.1)',
                tension: 0.3,
                fill: true
            }]
        },
        options: {
            responsive: true,
            maintainAspectRatio: false,
            scales: { y: { beginAtZero: true, max: 100 } },
            plugins: { legend: { display: false } }
        }
    });
}

function updateChart(chart, value) {
    const now = new Date().toLocaleTimeString();
    chart.data.labels.push(now);
    chart.data.datasets[0].data.push(value);
    if (chart.data.labels.length > 20) {
        chart.data.labels.shift();
        chart.data.datasets[0].data.shift();
    }
    chart.update();
}

function formatBytes(bytes) {
    if (bytes === 0) return '0 B';
    const k = 1024;
    const sizes = ['B', 'KB', 'MB', 'GB', 'TB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
}

function updateMetrics() {
    fetch('/api/metrics')
        .then(response => response.json())
        .then(data => {
            // CPU
            document.getElementById('cpu-cores').textContent = data.cpu.cores;
            document.getElementById('cpu-percent').textContent = data.cpu.percent.toFixed(1);
            updateChart(cpuChart, data.cpu.percent);

            // Memory
            document.getElementById('mem-total').textContent = (data.memory.total / 1e9).toFixed(1);
            document.getElementById('mem-used').textContent = (data.memory.used / 1e9).toFixed(1);
            document.getElementById('mem-percent').textContent = data.memory.used_percent.toFixed(1);
            updateChart(memChart, data.memory.used_percent);

            // Disk - генерируем карточки с прогресс-барами
            const diskDiv = document.getElementById('disk-stats');
            diskDiv.innerHTML = '';
            data.disk.forEach(disk => {
                const usedPercent = disk.used_percent.toFixed(1);
                const totalGB = (disk.total / 1e9).toFixed(1);
                const usedGB = (disk.used / 1e9).toFixed(1);
                const freeGB = (disk.free / 1e9).toFixed(1);
                const card = document.createElement('div');
                card.className = 'disk-item';
                card.innerHTML = `
                    <strong>${disk.mountpoint}</strong>
                    <p>${usedGB} GB / ${totalGB} GB (${usedPercent}%)</p>
                    <div class="progress-bar"><div class="progress-fill" style="width: ${usedPercent}%;"></div></div>
                    <small>Free: ${freeGB} GB</small>
                `;
                diskDiv.appendChild(card);
            });

            // Network
            const netDiv = document.getElementById('network-stats');
            netDiv.innerHTML = '';
            data.network.forEach(net => {
                const sent = formatBytes(net.bytes_sent);
                const recv = formatBytes(net.bytes_recv);
                const card = document.createElement('div');
                card.className = 'network-item';
                card.innerHTML = `
                    <strong>${net.interface_name}</strong>
                    <p>⬆️ ${sent} / ⬇️ ${recv}</p>
                `;
                netDiv.appendChild(card);
            });

            // Processes
            document.getElementById('procs-count').textContent = data.processes_count;
        })
        .catch(err => console.error('Error fetching metrics:', err));
}

window.onload = () => {
    initCharts();
    updateMetrics();
    setInterval(updateMetrics, 2000);
};