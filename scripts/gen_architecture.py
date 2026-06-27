#!/usr/bin/env python3
"""Generate an architecture overview diagram for the CVE Utils project."""

import matplotlib
import os
matplotlib.use('Agg')
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
from matplotlib.patches import FancyBboxPatch, FancyArrowPatch
import numpy as np

def draw_architecture():
    fig, ax = plt.subplots(1, 1, figsize=(22, 14))
    ax.set_xlim(-11, 11)
    ax.set_ylim(-7, 7)
    ax.set_aspect('equal')
    ax.axis('off')
    fig.patch.set_facecolor('#f0f0f5')

    def draw_box(x, y, w, h, label, sublabel, facecolor, edgecolor, fontsize=12, subfontsize=9, label_color='white', sub_color='#b2bec3'):
        box = FancyBboxPatch((x - w/2, y - h/2), w, h,
                              boxstyle="round,pad=0.15", facecolor=facecolor,
                              edgecolor=edgecolor, linewidth=2.5, zorder=5)
        ax.add_patch(box)
        ax.text(x, y + (0.12 if sublabel else 0), label, fontsize=fontsize, fontweight='bold',
                ha='center', va='center', zorder=6, color=label_color,
                fontfamily='sans-serif')
        if sublabel:
            ax.text(x, y - 0.25, sublabel, fontsize=subfontsize,
                    ha='center', va='center', zorder=6, style='italic', color=sub_color,
                    fontfamily='sans-serif')

    def draw_arrow(start, end, color='#636e72', lw=2):
        arrow = FancyArrowPatch(start, end, arrowstyle='->', color=color,
                                 linewidth=lw, connectionstyle="arc3,rad=0.1", zorder=3,
                                 mutation_scale=15)
        ax.add_patch(arrow)

    def draw_simple_arrow(start_x, start_y, end_x, end_y, color='#b2bec3', lw=1.5):
        ax.annotate('', xy=(end_x, end_y), xytext=(start_x, start_y),
                    arrowprops=dict(arrowstyle='->', color=color, lw=lw),
                    zorder=3)

    # === Title ===
    ax.text(0, 6.5, 'CVE Utils', fontsize=30, fontweight='bold',
            color='#2d3436', ha='center', va='center', fontfamily='sans-serif')
    ax.text(0, 5.9, 'Architecture Overview', fontsize=16,
            color='#636e72', ha='center', va='center', fontfamily='sans-serif')

    # === Layer 1: Application ===
    # CLI Tool
    draw_box(-4, 4.5, 3.8, 1.2, 'CLI Tool', 'cve command', '#2d3436', '#1a1a2e',
             fontsize=14, label_color='white')

    # Go Package (API)
    draw_box(4, 4.5, 3.8, 1.2, 'Go Package', 'github.com/scagogogo/cve-skills', '#e94560', '#c0392b',
             fontsize=14, label_color='white')

    draw_arrow((-2.1, 4.5), (2.1, 4.5), color='#e94560', lw=2.5)

    # === Layer 2: Modules ===
    modules = [
        (-8.5, 2.0, 2.8, 0.9, 'base.go', 'Format / Validate', '#ff6b6b', '#e55039'),
        (-4.8, 2.0, 2.8, 0.9, 'extract.go', 'Extract / Split', '#4ecdc4', '#00b894'),
        (-1.1, 2.0, 2.8, 0.9, 'compare.go', 'Compare / Sort', '#45b7d1', '#0984e3'),
        (2.6, 2.0, 2.8, 0.9, 'filter.go', 'Filter / Group', '#96ceb4', '#00b894'),
        (6.3, 2.0, 2.8, 0.9, 'generate.go', 'Generate / Range', '#ffeaa7', '#fdcb6e'),
    ]

    for mx, my, mw, mh, ml, ms, mc, me in modules:
        draw_box(mx, my, mw, mh, ml, ms, mc, me)
        draw_simple_arrow(4, 3.9, mx, my + mh/2, color='#b2bec3', lw=1.5)

    # === Layer 3: Function Categories ===
    # Organize as 3 rows, 5 columns (matching module positions)
    categories = [
        # Column 1 (base.go) - x=-8.5
        [(-8.5, -0.3, 'Format()', '#ff6b6b'),
         (-8.5, -1.0, 'IsCve()', '#ff6b6b'),
         (-8.5, -1.7, 'ValidateCve()', '#ff6b6b'),
         (-8.5, -2.4, 'IsCveYearOk()', '#ff6b6b'),
         (-8.5, -3.1, 'ValidateCves()', '#ff6b6b'),
         (-8.5, -3.8, 'FilterValidCves()', '#ff6b6b'),
         (-8.5, -4.5, 'IsContainsCve()', '#ff6b6b')],
        # Column 2 (extract.go) - x=-4.8
        [(-4.8, -0.3, 'ExtractCve()', '#4ecdc4'),
         (-4.8, -1.0, 'ExtractFirstCve()', '#4ecdc4'),
         (-4.8, -1.7, 'ExtractLastCve()', '#4ecdc4'),
         (-4.8, -2.4, 'Split()', '#4ecdc4'),
         (-4.8, -3.1, 'ExtractCveYear()', '#4ecdc4'),
         (-4.8, -3.8, 'ExtractCveSeq()', '#4ecdc4'),
         (-4.8, -4.5, '*AsInt variants', '#4ecdc4')],
        # Column 3 (compare.go) - x=-1.1
        [(-1.1, -0.3, 'CompareCves()', '#45b7d1'),
         (-1.1, -1.0, 'CompareByYear()', '#45b7d1'),
         (-1.1, -1.7, 'SubByYear()', '#45b7d1'),
         (-1.1, -2.4, 'SortCves()', '#45b7d1')],
        # Column 4 (filter.go) - x=2.6
        [(2.6, -0.3, 'FilterByYear()', '#96ceb4'),
         (2.6, -1.0, 'FilterByYearRange()', '#96ceb4'),
         (2.6, -1.7, 'GroupByYear()', '#96ceb4'),
         (2.6, -2.4, 'RemoveDuplicate()', '#96ceb4'),
         (2.6, -3.1, 'IntersectCves()', '#96ceb4'),
         (2.6, -3.8, 'UnionCves()', '#96ceb4'),
         (2.6, -4.5, 'DiffCves()', '#96ceb4')],
        # Column 5 (generate.go) - x=6.3
        [(6.3, -0.3, 'GenerateCve()', '#ffeaa7'),
         (6.3, -1.0, 'GenerateFakeCve()', '#ffeaa7'),
         (6.3, -1.7, 'FormatSeq()', '#ffeaa7'),
         (6.3, -2.4, 'ParseCveRange()', '#ffeaa7'),
         (6.3, -3.1, 'IsCvesConsecutive()', '#ffeaa7'),
         (6.3, -3.8, 'FilterByPattern()', '#ffeaa7'),
         (6.3, -4.5, 'CountByYear()', '#ffeaa7')],
    ]

    for col in categories:
        for cx, cy, cl, cc in col:
            text_w = max(len(cl) * 0.085 + 0.3, 0.9)
            ax.text(cx, cy, cl, fontsize=8.5, color='#2d3436', ha='center', va='center',
                    fontweight='bold', zorder=6, fontfamily='monospace',
                    bbox=dict(boxstyle='round,pad=0.25', facecolor='white', edgecolor=cc,
                              linewidth=1.5, alpha=0.92))

    # Module -> Category connectors (vertical lines)
    module_xs = [-8.5, -4.8, -1.1, 2.6, 6.3]
    for mx in module_xs:
        ax.plot([mx, mx], [1.55, -0.6], color='#b2bec3', linewidth=1, linestyle=':', alpha=0.5, zorder=1)

    # === Side: Infrastructure ===
    side_items = [
        (9.5, 4.5, 3.0, 0.7, 'Tests', '95%+ coverage', '#6c5ce7', '#5f27cd'),
        (9.5, 3.2, 3.0, 0.7, 'Examples', '30+ code samples', '#fd79a8', '#e84393'),
        (9.5, 1.9, 3.0, 0.7, 'Docs Site', 'VitePress', '#00cec9', '#00b894'),
        (9.5, 0.6, 3.0, 0.7, 'CI/CD', 'GitHub Actions', '#a29bfe', '#6c5ce7'),
        (9.5, -0.7, 3.0, 0.7, 'v0.0.1', 'MIT License', '#dfe6e9', '#b2bec3', '#2d3436', '#636e72'),
    ]
    for sx, sy, sw, sh, sl, ss, sc, se, *rest in side_items:
        lc = rest[0] if rest else 'white'
        sc2 = rest[1] if len(rest) > 1 else '#b2bec3'
        draw_box(sx, sy, sw, sh, sl, ss, sc, se, fontsize=10, label_color=lc, sub_color=sc2)

    # === Layer labels ===
    ax.text(-10.5, 5.0, 'Layer 1', fontsize=8, color='#b2bec3', rotation=90, va='center')
    ax.text(-10.5, 2.0, 'Layer 2', fontsize=8, color='#b2bec3', rotation=90, va='center')
    ax.text(-10.5, -2.5, 'Layer 3', fontsize=8, color='#b2bec3', rotation=90, va='center')

    # Horizontal divider lines
    ax.axhline(y=3.5, color='#b2bec3', linewidth=0.8, linestyle='--', alpha=0.4, xmin=0.02, xmax=0.85)
    ax.axhline(y=0.8, color='#b2bec3', linewidth=0.8, linestyle='--', alpha=0.4, xmin=0.02, xmax=0.85)

    plt.tight_layout(pad=0.5)
    plt.savefig(os.path.join(os.path.dirname(os.path.abspath(__file__)), '..', 'docs', 'images', 'architecture.png'),
                dpi=150, bbox_inches='tight', facecolor='#f0f0f5')
    print("Architecture diagram saved!")


if __name__ == '__main__':
    draw_architecture()
